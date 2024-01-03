package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	t.Parallel()

	actual := DefaultConfig()
	if assert.NotNil(t, actual, "should actual is not nil") {
		assert.Equal(t, 8080, actual.Port, "assert port is 8080")
		assert.False(t, actual.UseTLS, "assert use_tls is false")
		assert.Equal(t, "", actual.CrtPath, "assert crt_path is empty")
		assert.Equal(t, "", actual.KeyPath, "assert key_path is empty")
		assert.False(t, actual.CORS, "assert cors is false")
		assert.Equal(t, []string{}, actual.AllowOrigins, "assert allow_origins is empty")
	}
}

func TestConfig_Addr(t *testing.T) {
	t.Parallel()

	t.Run("should return empty string when config is nil", func(t *testing.T) {
		var c *Config
		assert.Empty(t, c.Addr(), "should return empty string")
	})

	t.Run("should return address when config is not nil", func(t *testing.T) {
		c := &Config{
			Port: 8080,
		}
		assert.Equal(t, ":8080", c.Addr(), "should return address")
	})
}
