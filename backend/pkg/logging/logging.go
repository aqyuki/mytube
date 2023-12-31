package logging

import (
	"log/slog"
	"time"

	"github.com/m-mizutani/clog"
)

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
