package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_ConnectionURL(t *testing.T) {
	t.Parallel()

	t.Run("returns empty string if config is nil", func(t *testing.T) {
		t.Parallel()

		var c *Config
		actual := c.ConnectionURL()
		assert.Empty(t, actual, "should return empty string, but got %s", actual)
	})

	t.Run("returns connection URL", func(t *testing.T) {
		t.Parallel()

		c := &Config{
			Name:     "name",
			User:     "user",
			Password: "password",
			Host:     "host",
			Port:     "port",
			SSLMode:  "sslmode",
		}

		excepted := "postgres://user:password@host:port/name?sslmode=sslmode"
		actual := c.ConnectionURL()
		assert.Equal(t, excepted, actual, "should return %s, but got %s", excepted, actual)
	})
}
