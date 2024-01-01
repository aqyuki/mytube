package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"log/slog"
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
	flag.Parse()

	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	logger := logging.NewLogger()
	ctx = logging.WithLogger(ctx, logger)

	defer func() {
		done()
	}()

	err := realMain(ctx)
	done()

	if err != nil {
		log.Fatal(err)
	}
	logger.Info("successful shutdown")
}

func realMain(ctx context.Context) error {
	logger := logging.FromContext(ctx)

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
