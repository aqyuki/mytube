package logging

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	t.Parallel()

	actual := NewLogger()
	assert.NotNil(t, actual, "NewLogger should not return nil, but received nil")
}

func TestContext(t *testing.T) {
	t.Parallel()

	t.Run("with logger", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		logger := NewLogger()

		actual := WithLogger(ctx, logger)
		assert.NotNil(t, actual, "WithLogger should not return nil, but received nil")
	})

	t.Run("nil logger", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		actual := FromContext(ctx)
		assert.NotNil(t, actual, "FromContext should not return nil, but received nil")
	})
}
