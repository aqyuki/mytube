package middleware

import (
	"log/slog"

	"github.com/aqyuki/mytube/backend/pkg/logging"
	"github.com/labstack/echo/v4"
)

func UnwrapLogger(c echo.Context) *slog.Logger {
	logger, ok := c.Get(contextLoggerKey).(*slog.Logger)
	if !ok {
		return logging.NewLogger()
	}
	return logger
}
