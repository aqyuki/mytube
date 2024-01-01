package setup

import (
	"context"
	"errors"

	"github.com/sethvargo/go-envconfig"
)

// Setup load configuration from environment variables
func Setup(ctx context.Context, config any) error {
	if config == nil {
		return errors.New("config is nil")
	}
	if err := envconfig.Process(ctx, config); err != nil {
		return err
	}
	return nil
}
