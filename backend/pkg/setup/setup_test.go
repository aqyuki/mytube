package setup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	type Env1 struct {
		A string `env:"A"`
		B string `env:"B"`
	}

	t.Run("should be success", func(t *testing.T) {
		t.Setenv("A", "a")
		t.Setenv("B", "b")
		t.Setenv("C", "c")
		t.Setenv("D", "d")

		var env Env1
		err := Setup(context.Background(), &env)
		assert.NoError(t, err, "should be success but got error")
		assert.Equal(t, env.A, "a", "should be equal")
		assert.Equal(t, env.B, "b", "should be equal")
	})

	t.Run("config is nil", func(t *testing.T) {
		err := Setup(context.Background(), nil)
		assert.Error(t, err, "should be error")
	})
}
