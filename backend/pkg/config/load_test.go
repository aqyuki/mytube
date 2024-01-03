package config

import (
	"context"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromEnv(t *testing.T) {
	t.Run("should return error when provided value is nil", func(t *testing.T) {
		err := LoadFromEnv(context.Background(), nil)
		assert.ErrorIs(t, err, ErrNilValue, "should return error when provided nil value")
	})

	t.Run("should return error when provided value is not a pointer", func(t *testing.T) {
		t.Setenv("test", "test")
		a := struct {
			A string `env:"test"`
		}{}

		err := LoadFromEnv(context.Background(), a)
		assert.ErrorIs(t, err, ErrNotPtr, "should return error when provided value is not a pointer")
	})

	t.Run("should return error when required environment is not set", func(t *testing.T) {
		a := struct {
			A string `env:"test,required"`
		}{}

		err := LoadFromEnv(context.Background(), &a)
		assert.ErrorIs(t, err, ErrNotRequired, "should return error when required environment is not set")
	})

	t.Run("success to load from environment variables", func(t *testing.T) {
		t.Setenv("test", "test")
		a := struct {
			A string `env:"test"`
		}{}

		err := LoadFromEnv(context.Background(), &a)
		assert.NoError(t, err, "should success to load from environment variables")
		assert.Equal(t, "test", a.A, "should success to load from environment variables")
	})
}

func TestLoadFromYaml(t *testing.T) {
	t.Parallel()

	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(filename), "testdata", "data.yml")

	t.Run("should return error when provided value is nil", func(t *testing.T) {
		t.Parallel()
		err := LoadFromYaml(context.Background(), path, nil)
		assert.ErrorIs(t, err, ErrNilValue, "should return error when provided nil value")
	})

	t.Run("should return error when file is not found", func(t *testing.T) {
		t.Parallel()
		err := LoadFromYaml(context.Background(), "not_found.yaml", &struct{}{})
		assert.ErrorIs(t, err, ErrNotFound, "should return error when file is not found")
	})

	t.Run("should return error when failed to open file", func(t *testing.T) {
		t.Parallel()
		err := LoadFromYaml(context.Background(), filepath.Join(filepath.Dir(filename), "testdata"), &struct{}{})
		assert.Error(t, err, "should return error when failed to open file")
	})

	t.Run("success to load from yaml", func(t *testing.T) {
		t.Parallel()
		a := struct {
			A string `yaml:"a"`
			B int    `yaml:"b"`
		}{}

		err := LoadFromYaml(context.Background(), path, &a)
		assert.NoError(t, err, "should success to load from yaml")
		assert.Equal(t, "hello", a.A, "should success to load from yaml")
		assert.Equal(t, 10, a.B, "should success to load from yaml")
	})
}
