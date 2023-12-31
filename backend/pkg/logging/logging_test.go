package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	t.Parallel()

	actual := NewLogger()
	assert.NotNil(t, actual, "NewLogger should not return nil, but received nil")
}
