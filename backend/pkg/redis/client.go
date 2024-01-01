package redis

import (
	"context"

	"github.com/aqyuki/mytube/backend/pkg/setup"
	"github.com/redis/go-redis/v9"
)

func NewClient(config *Config) (*redis.Client, error) {
	return redis.NewClient(config.Options()), nil
}

func NewClientFromEnv(ctx context.Context) (*redis.Client, error) {
	var config Config
	if err := setup.Setup(ctx, &config); err != nil {
		return nil, err
	}
	return NewClient(&config)
}
