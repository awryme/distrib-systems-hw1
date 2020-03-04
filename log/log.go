package log

import (
	"context"
	"go.uber.org/zap"
	"net/http"
)

type loggerCtxKey struct{}

func New() *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.DisableStacktrace = true
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger
}

func ToCtx(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerCtxKey{}, logger)
}

func FromCtx(ctx context.Context) *zap.Logger {
	ctxLogger := ctx.Value(loggerCtxKey{})
	if ctxLogger == nil {
		panic("no logger")
	}
	logger, ok := ctxLogger.(*zap.Logger)
	if !ok {
		panic("failed to cast logger")
	}
	return logger
}

func WithCtx(ctx context.Context, fields ...zap.Field) context.Context {
	logger := FromCtx(ctx)
	return ToCtx(ctx, logger.With(fields...))
}

func ToReq(r *http.Request, logger *zap.Logger) *http.Request {
	return r.WithContext(ToCtx(r.Context(), logger))
}

func FromReq(r *http.Request) *zap.Logger {
	return FromCtx(r.Context())
}

func WithReq(r *http.Request, fields ...zap.Field) *http.Request {
	return r.WithContext(WithCtx(r.Context(), fields...))
}
