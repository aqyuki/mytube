package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/aqyuki/mytube/backend/pkg/database"
	"github.com/aqyuki/mytube/backend/pkg/logging"
	"github.com/aqyuki/mytube/backend/pkg/setup"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	pathFlag = flag.String("path", "migrations/", "path to migrations folder")
)

func main() {
	if err := realMain(); err != nil {
		fmt.Printf("failed: %+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("server shutdown\n")
	os.Exit(0)
}

func realMain() error {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer done()

	logger := logging.NewLogger()

	var config database.Config
	if err := setup.Setup(ctx, &config); err != nil {
		return fmt.Errorf("failed database setup: %w", err)
	}

	dir := fmt.Sprintf("file://%s", *pathFlag)
	m, err := migrate.New(dir, config.ConnectionURL())
	if err != nil {
		return fmt.Errorf("failed create migrate: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed migrate up: %w", err)
	} else if errors.Is(err, migrate.ErrNoChange) {
		version, _, _ := m.Version()
		logger.Warn("database already migrated", slog.String("version", fmt.Sprintf("%v", version)))
	}

	srcErr, dbErr := m.Close()
	if srcErr != nil {
		return fmt.Errorf("migrate source error: %w", srcErr)
	}
	if dbErr != nil {
		return fmt.Errorf("migrate database error: %w", dbErr)
	}
	logger.Info("finished migrate")
	return nil
}

func init() {
	flag.Parse()
}
