package app

import (
	"context"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"speed/app/lib/log"
	"speed/app/lib/validate"
	"strings"
	"time"

	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/syyongx/php2go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

var (
	Config *viper.Viper
	Redis  *redis.Client
	Db     *gorm.DB
	Log    *zap.SugaredLogger
	Log0   *zap.Logger
	//var Log1 *zap.Logger
	AppName string
	AppPath string
	AppEnv  string
	AppKey  string
)

func InitApp(sqliteDbFile, config embed.FS) {

	initReadFsConfig(config)
	//initConfig()
	initAppName()

	initAppPath()
	initAppEnv()
	initAppKey()
	initLog()

	//initMysqlDb()/
	//initRedis()//
	initValidator()
	initSqliteDb(sqliteDbFile)

}

func initAppKey() {
	AppKey = viper.GetString("appKey")
}

func initValidator() {
	validate.Init()

}
func initAppEnv() {
	AppEnv = Config.GetString("appEnv")
}

func initConfig() {
	s, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(s)

	Config = viper.New()
	Config.SetConfigType("json")
	Config.SetConfigFile(".config.json")
	err := Config.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initReadFsConfig(configFS embed.FS) {

	Config = viper.New()
	Config.SetConfigType("json")
	f, err := configFS.Open(".config.json")
	if err != nil {
		panic(fmt.Errorf("failed to open embedded config: %w", err))
	}
	defer f.Close()

	if err := Config.ReadConfig(f); err != nil {
		panic(fmt.Errorf("failed to read config: %w", err))
	}
}

func initAppPath() {
	s, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	AppPath = php2go.Substr(s, 0, strings.Index(s, AppName)+len(AppName))
}

func initAppName() {
	AppName = Config.GetString("appName")
}

func initRedis() {
	var (
		host = Config.GetString("cache.redis.default.host")
		pass = Config.GetString("cache.redis.default.password")
		//idleConn, _ = Conf.Int("REDIS_MIN_IDLE")
	)

	Redis := redis.NewClient(&redis.Options{
		Addr:        host + ":6379",
		Password:    pass,             // Redis账号
		DB:          0,                // Redis库
		MaxRetries:  3,                // 最大重试次数
		IdleTimeout: 10 * time.Second, // 空闲链接超时时间
		Network:     "tcp",
	})
	pong, err := Redis.Ping().Result()
	if err == redis.Nil {
		log.Panic("Redis异常", err)
	} else if err != nil {
		log.Panic("redis异常:", err.Error())
	} else {
		log.Infof("redis init success %s ", pong)
	}

}

func initMysqlDb() {

	key := "db.mysql.default."
	var (
		db       *gorm.DB
		username = Config.GetString(key + "username")
		pass     = Config.GetString(key + "password")
		host     = Config.GetString(key + "host")
		port     = Config.GetString(key + "port")
		database = Config.GetString(key + "databaseName")
		charset  = Config.GetString(key + "charset")
		//dialect  = Config.GetString("db.dialect")
	)
	dsn := username + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database + "?charset=" + charset
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil || reflect.TypeOf(db).String() != "*gorm.DB" || db == nil {
		if err == nil || db == nil {
			log.Panicf("init DB connect failed db init false ")
			return
		}
		log.Panicf("init DB connect failed, error: %s", err.Error())
	}
	if err != nil {
		log.Panicf("init DB connect failed, error: %s", err.Error())
	}

	//if Config.GetString("appEnv") == "prod" {
	//	db.Config.Logger = Log
	//}

	Db = db
	log.Info("init DB connect success")

}

func initSqliteDb(sqliteFile embed.FS) {
	// 1. 确定一个可写的目标路径（例如用户的应用数据目录）
	homeDir, _ := os.UserHomeDir()
	dbDir := filepath.Join(homeDir, "Library", "Application Support", "wawa_shop", "data")
	os.MkdirAll(dbDir, 0755)
	targetPath := filepath.Join(dbDir, "shop.sqlite3")
	// 2. 从嵌入资源中读取数据
	data, _ := sqliteFile.ReadFile("data/shop.sqlite3")
	// 3. 将数据写入目标文件
	os.WriteFile(targetPath, data, 0644)
	// 4. 使用这个文件路径连接数据库

	Logger := zapgorm2.New(Log0)
	Logger.Context = func(ctx context.Context) []zapcore.Field {
		traceId := ctx.Value("traceId")
		return []zapcore.Field{{Key: "traceId", Type: zapcore.StringType, String: traceId.(string)}}
	}

	//targetPath = "data/shop.sqlite3"
	db, _ := gorm.Open(sqlite.Open(targetPath), &gorm.Config{Logger: Logger.LogMode(gormlogger.Info)})
	Db = db
}

func initLog() {
	Log0 = InitLog()
	Log = Log0.Sugar()
}

func InitLog() *zap.Logger {

	fileName := ""
	var level zapcore.Level
	if Config.Get("appEnv") == "prod" {
		fileName = AppPath + "/storage/logs/zap.log"
		level = getLoggerLevel("info")
	} else {
		fileName = "storage/logs/zap.log"
		level = getLoggerLevel("debug")
	}

	var write zapcore.WriteSyncer

	if AppEnv == "prod" {
		write = zapcore.AddSync(&lumberjack.Logger{
			Filename:  fileName,
			MaxSize:   1 << 30, //1G
			LocalTime: true,
			Compress:  true,
		})
	} else {
		write = os.Stdout
	}

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = func(i time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(i.Format("2006-01-02 15:04:05.000"))
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), write, zap.NewAtomicLevelAt(level))

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(log.Stack{}), zap.Development())
	log.Log = logger.Sugar()
	return logger
	//Log1 = logger //嵌套结构化
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}
