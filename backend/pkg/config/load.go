package config

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/sethvargo/go-envconfig"
	"gopkg.in/yaml.v3"
)

var (
	ErrNilValue    = errors.New("provided value is nil")
	ErrNotPtr      = errors.New("provided value is not a pointer")
	ErrNotRequired = errors.New("required environment is not set")
	ErrNotFound    = errors.New("not found file ")
)

// LoadFromEnv loads the configuration from the environment variables into the provided struct.
// The provided struct must be a pointer to a struct.
func LoadFromEnv(ctx context.Context, a any) error {
	if a == nil {
		return ErrNilValue
	}

	if err := envconfig.Process(ctx, a); err != nil {
		if errors.Is(err, envconfig.ErrNotPtr) {
			return ErrNotPtr
		} else if errors.Is(err, envconfig.ErrMissingRequired) {
			return ErrNotRequired
		}
		return fmt.Errorf("failed to load config from environment variables because %w", err)
	}
	return nil
}

// LoadFromYaml loads the configuration from the YAML file provided by the path into the provided struct.
// The provided struct must be a pointer to a struct.
func LoadFromYaml(ctx context.Context, path string, a any) error {
	if a == nil {
		return ErrNilValue
	}

	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return ErrNotFound
	} else if err != nil {
		return fmt.Errorf("failed to open file %s because %w", path, err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("failed to read file %s because %w", path, err)
	}

	if err := yaml.Unmarshal(b, a); err != nil {
		return fmt.Errorf("failed to unmarshal file %s because %w", path, err)
	}
	return nil
}
