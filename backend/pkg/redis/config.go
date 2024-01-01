package redis

import "github.com/redis/go-redis/v9"

type Config struct {
	Host     string `env:"REDIS_HOST, default=localhost"`
	Port     string `env:"REDIS_PORT, default=6379"`
	Password string `env:"REDIS_PASSWORD, default="`
	DB       int    `env:"REDIS_DB, default=0"`
	PoolSize int    `env:"REDIS_POOL_SIZE, default=1000"`
}

func (c *Config) Options() *redis.Options {
	return &redis.Options{
		Addr:     c.Host + ":" + c.Port,
		Password: c.Password,
		DB:       c.DB,
		PoolSize: c.PoolSize,
	}
}
