package logs

import (
	"context"
	"fmt"
)

func Warnf(ctx context.Context, format string, v ...any) {
	OtLogger.Ctx(ctx).Warn(fmt.Sprintf(format, v...))
}

func Infof(ctx context.Context, format string, v ...any) {
	OtLogger.Ctx(ctx).Info(fmt.Sprintf(format, v...))
}

func Errorf(ctx context.Context, format string, v ...any) {
	OtLogger.Ctx(ctx).Error(fmt.Sprintf(format, v...))
}

func Debugf(ctx context.Context, format string, v ...any) {
	OtLogger.Ctx(ctx).Debug(fmt.Sprintf(format, v...))
}

func Panicf(ctx context.Context, format string, v ...any) {
	OtLogger.Ctx(ctx).Panic(fmt.Sprintf(format, v...))
}

func Fatalf(ctx context.Context, format string, v ...any) {
	OtLogger.Ctx(ctx).Fatal(fmt.Sprintf(format, v...))
}
