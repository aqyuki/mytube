package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aqyuki/mytube/backend/pkg/account"
	"github.com/aqyuki/mytube/backend/pkg/database"
	"github.com/aqyuki/mytube/backend/pkg/logging"
	"github.com/aqyuki/mytube/backend/pkg/redis"
	"github.com/aqyuki/mytube/backend/pkg/server"
	"github.com/aqyuki/mytube/backend/pkg/session"
	"github.com/aqyuki/mytube/backend/pkg/setup"
	"github.com/joho/godotenv"
	"github.com/rbcervilla/redisstore/v9"
)

var (
	envPath = flag.String("env", "", "path to .env file")
)

func main() {
	if err := preMain(); err != nil {
		fmt.Printf("failed: %v\n", err)
		os.Exit(1)
	}
	if err := realMain(); err != nil {
		fmt.Printf("failed: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func preMain() error {
	if *envPath == "" {
		return nil
	}

	fmt.Printf("loading env from %s\n", *envPath)
	if err := godotenv.Load(*envPath); err != nil {
		return fmt.Errorf("failed load env: %w", err)
	}
	return nil
}

func realMain() error {
	logger := logging.NewLogger()

	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer done()

	ctx = logging.WithLogger(ctx, logger)

	// create database connection
	dbConn, err := database.NewConnFromEnv(ctx)
	if err != nil {
		return fmt.Errorf("failed create database connection: %w", err)
	}
	defer dbConn.Close()
	logger.Info("database connection created")

	// create redis connection
	redisConn, err := redis.NewClientFromEnv(ctx)
	if err != nil {
		return fmt.Errorf("failed create redis connection: %w", err)
	}
	defer redisConn.Close()
	logger.Info("redis connection created")

	// create redis store to store session
	store, err := redisstore.NewRedisStore(ctx, redisConn)
	if err != nil {
		return fmt.Errorf("failed create redis store: %w", err)
	}

	// initialize services
	accountService := account.NewAccountService(account.NewAccountRepository(dbConn))
	sessionManager := session.NewManager(store)

	// initialize modules
	modules := server.Modules{
		AccountService: accountService,
		SessionManager: sessionManager,
	}

	logger.Info("server initialized")

	var config server.Config
	if err := setup.Setup(ctx, &config); err != nil {
		return fmt.Errorf("failed setup: %w", err)
	}
	server := server.New(ctx, &modules, &config)

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Error("failed shutdown server: %v", err)
			return
		}
		logger.Info("server shutdown")
	}()

	logger.Info("server start. Waiting for request...")
	if err := server.Start(8080); err != nil {
		return fmt.Errorf("failed listen and serve: %w", err)
	}
	return nil
}

func init() {
	flag.Parse()
}
