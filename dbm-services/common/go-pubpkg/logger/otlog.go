package logger

import (
	"context"
	"fmt"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap/zapcore"
)

var OtLogger = otelzap.New(GetLogger(),
	otelzap.WithTraceIDField(true),
	otelzap.WithCaller(true),
	otelzap.WithStackTrace(true),
	otelzap.WithMinLevel(zapcore.InfoLevel),
	otelzap.WithErrorStatusLevel(zapcore.ErrorLevel),
)

// Warnf TODO
func Warnf(ctx context.Context, format string, v ...any) {
	OtLogger.Ctx(ctx).Warn(fmt.Sprintf(format, v...))
}

// Infof TODO
func Infof(ctx context.Context, format string, v ...any) {
	OtLogger.Ctx(ctx).Info(fmt.Sprintf(format, v...))
}

// Errorf TODO
func Errorf(ctx context.Context, format string, v ...any) {
	OtLogger.Ctx(ctx).Error(fmt.Sprintf(format, v...))
}

// Debugf TODO
func Debugf(ctx context.Context, format string, v ...any) {
	OtLogger.Ctx(ctx).Debug(fmt.Sprintf(format, v...))
}

// Panicf TODO
func Panicf(ctx context.Context, format string, v ...any) {
	OtLogger.Ctx(ctx).Panic(fmt.Sprintf(format, v...))
}

// Fatalf TODO
func Fatalf(ctx context.Context, format string, v ...any) {
	OtLogger.Ctx(ctx).Fatal(fmt.Sprintf(format, v...))
}
