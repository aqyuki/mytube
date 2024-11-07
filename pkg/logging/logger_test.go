package logging

import (
	"context"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewLoggerFromEnv(t *testing.T) {
	t.Run("When LOG_MODE is develop", func(t *testing.T) {
		t.Setenv("LOG_MODE", "develop")
		actual := NewLoggerFromEnv()

		if actual == nil {
			t.Errorf("Expected logger to be not nil, but got nil")
		}
	})

	t.Run("When LOG_MODE is not develop", func(t *testing.T) {
		t.Setenv("LOG_MODE", "production")
		actual := NewLoggerFromEnv()

		if actual == nil {
			t.Errorf("Expected logger to be not nil, but got nil")
		}
	})
}

func TestNewLogger(t *testing.T) {
	t.Parallel()
	t.Run("When develop is true", func(t *testing.T) {
		t.Parallel()
		actual := NewLogger(true, "debug")

		if actual == nil {
			t.Errorf("Expected logger to be not nil, but got nil")
		}
	})

	t.Run("When develop is false", func(t *testing.T) {
		t.Parallel()
		actual := NewLogger(false, "info")

		if actual == nil {
			t.Errorf("Expected logger to be not nil, but got nil")
		}
	})
}

func TestDefaultLogger(t *testing.T) {
	t.Parallel()

	actual1 := DefaultLogger()
	actual2 := DefaultLogger()

	if actual1 == nil || actual2 == nil {
		t.Errorf("Expected both loggers to be not nil")
	}
	if actual1 != actual2 {
		t.Errorf("Expected both loggers to be the same, but got different instances")
	}
}

func TestWithLogger(t *testing.T) {
	t.Parallel()
	logger := zap.NewNop()
	ctx := WithLogger(context.Background(), logger)

	if ctx == nil {
		t.Errorf("Expected context with logger to be not nil, but got nil")
	}
}

func TestFromContext(t *testing.T) {
	t.Parallel()

	t.Run("When logger is set in context", func(t *testing.T) {
		t.Parallel()
		logger := zap.NewNop()
		ctx := WithLogger(context.Background(), logger)
		actual := FromContext(ctx)

		if actual == nil {
			t.Errorf("Expected logger from context to be not nil, but got nil")
		}
	})

	t.Run("When logger is not set in context", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		actual := FromContext(ctx)

		if actual == nil {
			t.Errorf("Expected default logger to be returned, but got nil")
		}
	})
}

func Test_levelToZapLevel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		level string
		want  zapcore.Level
	}{
		{"debug", "debug", zapcore.DebugLevel},
		{"info", "info", zapcore.InfoLevel},
		{"warning", "warning", zapcore.WarnLevel},
		{"error", "error", zapcore.ErrorLevel},
		{"critical", "critical", zapcore.DPanicLevel},
		{"alert", "alert", zapcore.PanicLevel},
		{"emergency", "emergency", zapcore.FatalLevel},
		{"unknown", "unknown", zapcore.InfoLevel},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := levelToZapLevel(tt.level)
			if actual != tt.want {
				t.Errorf("Expected level %v, but got %v", tt.want, actual)
			}
		})
	}
}

func Test_levelEncoder(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		level zapcore.Level
		want  string
	}{
		{"debug", zapcore.DebugLevel, "DEBUG"},
		{"info", zapcore.InfoLevel, "INFO"},
		{"warning", zapcore.WarnLevel, "WARNING"},
		{"error", zapcore.ErrorLevel, "ERROR"},
		{"critical", zapcore.DPanicLevel, "CRITICAL"},
		{"alert", zapcore.PanicLevel, "ALERT"},
		{"emergency", zapcore.FatalLevel, "EMERGENCY"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mem := zapcore.NewMapObjectEncoder()
			err := mem.AddArray("k", zapcore.ArrayMarshalerFunc(func(arr zapcore.ArrayEncoder) error {
				levelEncoder()(tt.level, arr)
				return nil
			}))
			if err != nil {
				t.Errorf("Expected no error, but got %v", err)
			}
			arr := mem.Fields["k"].([]any)
			if len(arr) != 1 {
				t.Errorf("Expected array length 1, but got %d", len(arr))
			}
			if arr[0] != tt.want {
				t.Errorf("Expected %v, but got %v", tt.want, arr[0])
			}
		})
	}
}

func Test_timeEncoder(t *testing.T) {
	t.Parallel()
	moment := time.Unix(100, 50005000).UTC()
	mem := zapcore.NewMapObjectEncoder()
	err := mem.AddArray("k", zapcore.ArrayMarshalerFunc(func(arr zapcore.ArrayEncoder) error {
		timeEncoder()(moment, arr)
		return nil
	}))
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	arr := mem.Fields["k"].([]any)
	if len(arr) != 1 {
		t.Errorf("Expected array length 1, but got %d", len(arr))
	}
	expected := "1970-01-01T00:01:40.050005Z"
	if arr[0] != expected {
		t.Errorf("Expected %v, but got %v", expected, arr[0])
	}
}
