package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Options(t *testing.T) {
	t.Parallel()

	config := Config{
		Host:     "localhost",
		Port:     "6379",
		Password: "",
		DB:       0,
		PoolSize: 1000,
	}

	actual := config.Options()
	if assert.NotNil(t, actual) {
		assert.Equal(t, "localhost:6379", actual.Addr)
		assert.Equal(t, "", actual.Password)
		assert.Equal(t, 0, actual.DB)
		assert.Equal(t, 1000, actual.PoolSize)
	}
}
