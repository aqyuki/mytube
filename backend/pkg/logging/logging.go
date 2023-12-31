package logging

import (
	"context"
	"log/slog"
	"time"

	"github.com/m-mizutani/clog"
)

type contextKey string

const loggerKey contextKey = "logger"

// NewLogger creates a new slog.Logger
func NewLogger() *slog.Logger {
	handler := clog.New(
		clog.WithColor(true),
		clog.WithLevel(slog.LevelInfo),
		clog.WithTimeFmt(time.DateTime),
		clog.WithPrinter(clog.LinearPrinter),
	)
	return slog.New(handler)
}

// WithLogger creates a new context with a slog.Logger
// WithLogger returned panic when ctx is nil
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext returns a slog.Logger from context
// If not exist logger in context, return NewLogger()
func FromContext(ctx context.Context) *slog.Logger {
	l, ok := ctx.Value(loggerKey).(*slog.Logger)
	if ok {
		return l
	}
	return NewLogger()
}
