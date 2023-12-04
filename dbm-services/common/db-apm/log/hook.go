package logs

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once sync.Once
)

// setDefaultConfig
func setDefaultConfig() {
	viper.SetDefault(LevelConfigPath, "info")
	viper.SetDefault(PathConfigPath, "")
}

// 从string到AtomicLevel的转换
func setLogLevel(level string) {
	switch level {
	case "debug":
		LoggerLevel.SetLevel(zap.DebugLevel)
	case "info":
		LoggerLevel.SetLevel(zap.InfoLevel)
	case "warning":
		LoggerLevel.SetLevel(zap.WarnLevel)
	case "error":
		LoggerLevel.SetLevel(zap.ErrorLevel)
	case "fatal":
		LoggerLevel.SetLevel(zap.FatalLevel)
	default:
		LoggerLevel.SetLevel(zap.InfoLevel)
	}
}

// 初始化日志配置
func initLogConfig() {
	var (
		encoder zapcore.Encoder
		err     error
	)

	// 配置日志级别
	setLogLevel(viper.GetString(LevelConfigPath))

	// 日志路径及轮转配置
	var writeSyncer zapcore.WriteSyncer
	if viper.GetString(PathConfigPath) == "" {
		writeSyncer = zapcore.Lock(os.Stdout)
	} else {
		if Syncer, err = NewReopenableWriteSyncer(viper.GetString(PathConfigPath)); err != nil {
			fmt.Printf("failed to create syncer for->[%s]", err)
			return
		}
		writeSyncer = Syncer
	}
	// 配置日志格式
	encoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	ZapLogger = zap.New(
		zapcore.NewCore(encoder, writeSyncer, LoggerLevel),
		zap.AddCaller(), zap.AddCallerSkip(1),
	)

	// 追加两个option：调用来源、Error级别增加调用栈
	OtLogger = otelzap.New(ZapLogger,
		otelzap.WithTraceIDField(true),
		otelzap.WithCaller(true),
		otelzap.WithStackTrace(true),
		otelzap.WithMinLevel(zapcore.InfoLevel),
		otelzap.WithErrorStatusLevel(zapcore.ErrorLevel),
	)

	if OtLogger == nil {
		fmt.Printf("failed to build logger for it still is nil, log may disabled.")
		return
	}

	OtLogger.Info("logger config success.I am on duty now!")
}

// init
func init() {
	fmt.Println("logs hook init.")
	once.Do(func() {
		setDefaultConfig()
		initLogConfig()
	})
}

// InitTestLogger 加载单元测试日志配置
func InitTestLogger() {
	// 加载配置
	once.Do(func() {
		viper.Set(LevelConfigPath, "debug")
		initLogConfig()
	})
}
