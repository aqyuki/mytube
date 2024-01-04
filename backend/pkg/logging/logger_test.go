package logging

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	t.Parallel()

	actual := NewLogger()
	assert.NotNil(t, actual, "NewLogger should not return nil")
}

func TestContextWithLogger(t *testing.T) {
	t.Parallel()

	ctx := ContextWithLogger(context.Background(), nil)
	assert.NotNil(t, ctx, "ContextWithLogger should not return nil")
}

func TestLoggerFromContext(t *testing.T) {
	t.Parallel()

	actual := LoggerFromContext(context.Background())
	assert.NotNil(t, actual, "LoggerFromContext should not return nil")
}
