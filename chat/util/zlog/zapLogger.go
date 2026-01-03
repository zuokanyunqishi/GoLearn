package zlog

import (
	"GoLearn/chat/util/color"
	timeFormat "GoLearn/chat/util/format/time"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"strings"
	"time"
)

var zapLogger *zap.Logger

func Init(logFile string) {
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    500, // megabytes
		MaxBackups: 0,
		MaxAge:     28, // days
		LocalTime:  true,
	})
	sync := zapcore.AddSync(writer)
	jsonEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
	core := zapcore.NewTee(zapcore.NewCore(jsonEncoder, sync, infoLevel))
	zapLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer zapLogger.Sync()
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Info(strings.Repeat("-", 80))
}

func PrintInfo(msg string) {
	fmt.Printf(color.Green(time.Now().Format(timeFormat.YmdHis)+" #%s\n"), msg)
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Info(msg)
}

func PrintFatal(msg string) {
	fmt.Printf(color.Red(time.Now().Format(timeFormat.YmdHis)+" #%s\n"), msg)
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Fatal(msg)
}

func PrintInfof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf(color.Green(time.Now().Format(timeFormat.YmdHis)+" #%s\n"), msg)
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Info(msg)
}

func PrintErrorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf(color.Red(time.Now().Format(timeFormat.YmdHis)+" #%s\n"), msg)
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Error(msg)
}

func PrintFatalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf(color.Red(time.Now().Format(timeFormat.YmdHis)+" #%s\n"), msg)
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Fatal(msg)
}

func PrintWarnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf(color.Yellow(time.Now().Format(timeFormat.YmdHis)+" #%s\n"), msg)
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Warn(msg)
}

func PrintError(msg string) {
	fmt.Printf(color.Red(time.Now().Format(timeFormat.YmdHis)+" #%s\n"), msg)
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Error(msg)
}

func Info(msg string) {
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Info(msg)
}

func Warn(msg string) {
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Warn(msg)
}
func Fatal(msg string) {
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Fatal(msg)

}

func Error(msg string) {
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Error(msg)
}

func Infof(format string, args ...interface{}) {
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Infof(format, args...)

}
func Fatalf(format string, args ...interface{}) {
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Fatalf(format, args...)

}

func Errorf(format string, args ...interface{}) {
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Errorf(format, args...)

}
func Warnf(format string, args ...interface{}) {
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Warnf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	zapLogger.Sugar().With("format", time.Now().Format(timeFormat.YmdHis)).Debugf(format, args...)
}
