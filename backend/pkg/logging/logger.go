package logging

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/m-mizutani/clog"
)

type contextKey string

var (
	defaultLogger     *slog.Logger
	defaultLoggerOnce sync.Once
)

const (
	loggerKey contextKey = "logger"
)

func DefaultLogger() *slog.Logger {
	defaultLoggerOnce.Do(func() {
		defaultLogger = NewLogger()
	})
	return defaultLogger
}

func NewLogger() *slog.Logger {
	return slog.New(
		clog.New(
			clog.WithColor(true),
			clog.WithLevel(slog.LevelInfo),
			clog.WithPrinter(clog.LinearPrinter),
			clog.WithTimeFmt(time.DateTime),
		))
}

func ContextWithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func LoggerFromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		return logger
	}
	return DefaultLogger()
}
