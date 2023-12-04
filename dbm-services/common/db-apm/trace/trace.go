package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

const (
	// TracerName trace name
	TracerName = "bk-dbm/db-apm"
)

// InsertIntIntoSpan
func InsertIntIntoSpan(key string, value int, span oteltrace.Span) {
	if span == nil {
		return
	}
	span.SetAttributes(
		attribute.Int(key, value),
	)
}

// InsertStringIntoSpan
func InsertStringIntoSpan(key, value string, span oteltrace.Span) {
	if span == nil {
		return
	}
	span.SetAttributes(
		attribute.String(key, value),
	)
}

// InsertStringSliceIntoSpan
func InsertStringSliceIntoSpan(key string, value []string, span oteltrace.Span) {
	if span == nil {
		return
	}
	span.SetAttributes(
		attribute.StringSlice(key, value),
	)
}

// InsertIntSliceIntoSpan
func InsertIntSliceIntoSpan(key string, value []int, span oteltrace.Span) {
	if span == nil {
		return
	}
	span.SetAttributes(
		attribute.IntSlice(key, value),
	)
}

// IntoContext 填充trace，并返回处理后的context和span
// span为nil时，说明没有开启trace
func IntoContext(globalCtx context.Context, tracerName, spanName string) (context.Context, oteltrace.Span) {
	var (
		span     oteltrace.Span
		traceCtx context.Context
	)

	// 向trace context中添加trace
	tracer := otel.Tracer(tracerName)
	traceCtx, span = tracer.Start(globalCtx, spanName)

	return traceCtx, span

}
