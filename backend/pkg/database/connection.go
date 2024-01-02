package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aqyuki/mytube/backend/pkg/setup"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// NewConn creates a new postgres connection.
func NewConn(c *Config) (*sql.DB, error) {
	return sql.Open("pgx", c.ConnectionURL())
}

// NewConnFromEnv creates a new postgres connection from environment variables.
func NewConnFromEnv(ctx context.Context) (*sql.DB, error) {
	var c Config
	if err := setup.Setup(ctx, &c); err != nil {
		return nil, fmt.Errorf("failed database setup: %w", err)
	}
	return NewConn(&c)
}
