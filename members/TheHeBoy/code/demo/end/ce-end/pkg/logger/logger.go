// Package logger 处理日志相关逻辑
package logger

import (
	"fmt"
	"gohub/pkg/app"
	"gohub/pkg/config"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger 全局 Logger 对象
var Logger *zap.SugaredLogger
var LogZap *zap.Logger

// InitLogger 日志初始化
func InitLogger() {
	filename := config.GetString("log.filename")
	maxSize := config.GetInt("log.max_size")
	maxBackup := config.GetInt("log.max_backup")
	maxAge := config.GetInt("log.max_age")
	compress := config.GetBool("log.compress")
	logType := config.GetString("log.type")
	level := config.GetString("log.level")
	// 获取日志写入介质
	writeSyncer := getLogWriter(filename, maxSize, maxBackup, maxAge, compress, logType)

	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		panic(err)
	}

	// 初始化 core
	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)

	// 初始化 Logger
	LogZap = zap.New(core,
		zap.AddCaller(),      // 调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1), // 封装了一层，调用文件去除一层(runtime.Caller(1))
	)
	Logger = LogZap.Sugar()

	defer Logger.Sync()
}

// getEncoder 设置日志存储格式
func getEncoder() zapcore.Encoder {

	// 日志格式规则
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller", // 代码调用，如 page/page.go:148
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志级别名称大写，如 ERROR、INFO
		EncodeTime:     customTimeEncoder,              // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}

	// 本地环境配置
	if app.IsLocal() {
		// 终端输出的关键词高亮
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		// 本地设置内置的 Console 解码器（支持 stacktrace 换行）
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 线上环境使用 JSON 编码器
	return zapcore.NewJSONEncoder(encoderConfig)
}

// customTimeEncoder 自定义友好的时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// getLogWriter 日志记录介质。Gohub 中使用了两种介质，os.Stdout 和文件
func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string) zapcore.WriteSyncer {

	// 如果配置了按照日期记录日志文件
	if logType == "daily" {
		logname := time.Now().Format("2006-01-02.log")
		filename = strings.ReplaceAll(filename, "logs.log", logname)
	}

	// 滚动日志，详见 config
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
}

// Debug is a package-level function that calls the Debug method on the Logger
func Debug(args ...any) {
	Logger.Debug(args)
}

// Info is a package-level function that calls the Info method on the Logger
func Info(args ...any) {
	Logger.Info(args)
}

// Warn is a package-level function that calls the Warn method on the Logger
func Warn(args ...any) {
	Logger.Warn(args)
}

// Error is a package-level function that calls the Error method on the Logger
func Error(args ...any) {
	Logger.Error(args)
}

func Errorv(err error) {
	_, ok := err.(fmt.Formatter)
	if ok {
		Logger.Errorf("%+v", err)
	} else {
		Logger.Error(err)
	}
}

func ErrorIf(err error) {
	if err != nil {
		Logger.Error(err)
	}
}

// Debugw is a package-level function that calls the Debugw method on the Logger
func Debugw(msg string, keysAndValues ...interface{}) {
	Logger.Debugw(msg, keysAndValues...)
}

// Infow is a package-level function that calls the Infow method on the Logger
func Infow(msg string, keysAndValues ...interface{}) {
	Logger.Infow(msg, keysAndValues...)
}

// Warnw is a package-level function that calls the Warnw method on the Logger
func Warnw(msg string, keysAndValues ...interface{}) {
	Logger.Warnw(msg, keysAndValues...)
}

// Errorw is a package-level function that calls the Errorw method on the Logger
func Errorw(msg string, keysAndValues ...interface{}) {
	Logger.Errorw(msg, keysAndValues...)
}

// DPanic is a package-level function that calls the DPanic method on the Logger
func DPanic(args ...any) {
	Logger.DPanic(args...)
}

// Panic is a package-level function that calls the Panic method on the Logger
func Panic(args ...any) {
	Logger.Panic(args...)
}

// Fatal is a package-level function that calls the Fatal method on the Logger
func Fatal(args ...any) {
	Logger.Fatal(args...)
}

// Debugf is a package-level function that calls the Debugf method on the Logger
func Debugf(template string, args ...any) {
	Logger.Debugf(template, args...)
}

// Infof is a package-level function that calls the Infof method on the Logger
func Infof(template string, args ...any) {
	Logger.Infof(template, args...)
}

// Warnf is a package-level function that calls the Warnf method on the Logger
func Warnf(template string, args ...any) {
	Logger.Warnf(template, args...)
}

// Errorf is a package-level function that calls the Errorf method on the Logger
func Errorf(template string, args ...any) {
	Logger.Errorf(template, args...)
}

// DPanicf is a package-level function that calls the DPanicf method on the Logger
func DPanicf(template string, args ...any) {
	Logger.DPanicf(template, args...)
}

// Panicf is a package-level function that calls the Panicf method on the Logger
func Panicf(template string, args ...any) {
	Logger.Panicf(template, args...)
}

// Fatalf is a package-level function that calls the Fatalf method on the Logger
func Fatalf(template string, args ...any) {
	Logger.Fatalf(template, args...)
}
