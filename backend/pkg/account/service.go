package account

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/aqyuki/mytube/backend/pkg/logging"
	"github.com/rs/xid"
)

// Service is the interface that provides account methods.
type Service interface {
	// Register registers a new account
	Register(ctx context.Context, username, password string) (*Account, error)
	// Delete deletes an account
	Delete(ctx context.Context, username string) error
	// Login authenticates an account
	Login(ctx context.Context, username, password string) (*Account, error)
}

// AccountService is the implementation of the Service interface.
type AccountService struct {
	repo Repository
}

func NewAccountService(repo Repository) *AccountService {
	return &AccountService{repo: repo}
}

// Register registers a account to the repository
func (s *AccountService) Register(ctx context.Context, username, password string) (*Account, error) {
	logger := logging.FromContext(ctx)
	logger.Info("registering account")

	user, err := NewAccount(xid.New().String(), username, password, time.Now(), time.Now())
	if err != nil {
		logger.Error("failed to create account", slog.Any("error", err))
		return nil, err
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()

	if err := s.repo.Register(ctx, user); err != nil {
		logger.Error("failed to register account", slog.Any("error", err))
		return nil, err
	}

	logger.Info("account registered")
	return user, nil
}

// Delete deletes an account from the repository
func (s *AccountService) Delete(ctx context.Context, username string) error {
	logger := logging.FromContext(ctx)
	logger.Info("deleting account")

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
	defer cancel()

	user, err := s.repo.FetchByUsername(ctx, username)
	if err != nil {
		logger.Error("failed to fetch account", slog.Any("error", err))
		return err
	}
	if err := s.repo.Delete(ctx, user); err != nil {
		logger.Error("failed to delete account", slog.Any("error", err))
		return err
	}

	logger.Info("account deleted")
	return nil
}

func (s *AccountService) Login(ctx context.Context, username, password string) (*Account, error) {
	logger := logging.FromContext(ctx)
	logger.Info("logging in")

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
	defer cancel()

	user, err := s.repo.FetchByUsername(ctx, username)
	if err != nil {
		logger.Error("failed to fetch account", slog.Any("error", err))
		return nil, err
	}

	if user.EqualPassword(password) {
		return user, nil
	}
	return nil, errors.New("failed to login")
}
