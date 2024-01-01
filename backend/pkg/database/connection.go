package database

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sethvargo/go-envconfig"
)

// NewConn creates a new postgres connection.
func NewConn(c *Config) (*sql.DB, error) {
	return sql.Open("pgx", c.ConnectionURL())
}

// NewConnFromEnv creates a new postgres connection from environment variables.
func NewConnFromEnv(ctx context.Context) (*sql.DB, error) {
	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		return nil, errors.New("failed to load environment variables")
	}
	return NewConn(&c)
}
