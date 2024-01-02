package middleware

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

const contextLoggerKey = "echo-context-logger"

func NewLogger(logger *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger.Info(
				"Request received",
				slog.String("method", c.Request().Method),
				slog.String("path", c.Request().URL.Path),
				slog.String("ip", c.RealIP()),
				slog.String("user-agent", c.Request().UserAgent()),
			)
			return next(c)
		}
	}
}

func NewStoreLogger(logger *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(contextLoggerKey, logger)
			return next(c)
		}
	}
}
