package redis

import "github.com/redis/go-redis/v9"

type Config struct {
	Address  string `env:"REDIS_ADDRESS, default=localhost:6379"`
	Password string `env:"REDIS_PASSWORD"`
	DB       int    `env:"REDIS_DB, default=0"`
	PoolSize int    `env:"REDIS_POOL_SIZE, default=1000"`
}

func (c *Config) RedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     c.Address,
		Password: c.Password,
		DB:       c.DB,
		PoolSize: c.PoolSize,
	}
}
