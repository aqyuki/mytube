package redis

import (
	"context"

	"github.com/aqyuki/mytube/backend/pkg/setup"
	"github.com/redis/go-redis/v9"
)

func NewRedisConn(ctx context.Context) (*redis.Client, error) {
	var config Config
	if err := setup.Setup(ctx, &config); err != nil {
		return nil, err
	}
	return redis.NewClient(config.RedisOptions()), nil
}
