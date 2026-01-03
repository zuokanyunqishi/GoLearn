package app

import (
	"context"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
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
	//"modernc.org/sqlite" // 使用纯 Go 实现的 SQLite 驱动，不需要 CGO
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
	// 1. 确定一个可写的目标路径（跨平台支持）
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Panicf("无法获取用户主目录: %v", err)
	}

	// 根据操作系统选择合适的数据目录
	var dbDir string
	if AppEnv == "prod" {
		// 生产环境：使用应用数据目录
		if runtime.GOOS == "windows" {
			// Windows: 使用 AppData\Local
			dbDir = filepath.Join(homeDir, "AppData", "Local", AppName, "data")
		} else if runtime.GOOS == "darwin" {
			// macOS: 使用 Library/Application Support
			dbDir = filepath.Join(homeDir, "Library", "Application Support", AppName, "data")
		} else {
			// Linux 和其他系统: 使用 .config
			dbDir = filepath.Join(homeDir, ".config", AppName, "data")
		}
	} else {
		// 开发环境：使用项目目录
		dbDir = filepath.Join(AppPath, "data")
	}

	if err := os.MkdirAll(dbDir, 0755); err != nil {
		log.Panicf("无法创建数据库目录: %v", err)
	}

	targetPath := filepath.Join(dbDir, "shop.sqlite3")

	// 2. 从嵌入资源中读取数据
	data, err := sqliteFile.ReadFile("data/shop.sqlite3")
	if err != nil {
		log.Panicf("无法读取嵌入的数据库文件: %v", err)
	}

	// 3. 将数据写入目标文件（如果文件不存在才写入）
	// if _, err := os.Stat(targetPath); os.IsNotExist(err) {
	if err := os.WriteFile(targetPath, data, 0644); err != nil {
		log.Panicf("无法写入数据库文件: %v", err)
	}
	// }

	// 4. 使用这个文件路径连接数据库
	Logger := zapgorm2.New(Log0)
	Logger.Context = func(ctx context.Context) []zapcore.Field {
		traceId := ctx.Value("traceId")
		if traceId != nil {
			if str, ok := traceId.(string); ok {
				return []zapcore.Field{{Key: "traceId", Type: zapcore.StringType, String: str}}
			}
		}
		return []zapcore.Field{}
	}

	db, err := gorm.Open(sqlite.Open(targetPath), &gorm.Config{Logger: Logger.LogMode(gormlogger.Info)})
	if err != nil {
		log.Panicf("无法连接数据库: %v", err)
	}
	if db == nil {
		log.Panic("数据库连接初始化失败: db 为 nil")
	}

	Db = db
	log.Info("数据库初始化成功")
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
