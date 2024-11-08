package account

import (
	"context"
	"errors"

	"github.com/aqyuki/mytube/pkg/encrypt"
	"github.com/aqyuki/mytube/pkg/identifier"
	"github.com/samber/oops"
)

var (
	ErrUsernameConflict     = errors.New("the username is not usable because it is already used")
	ErrAuthenticationFailed = errors.New("the username or password is not correct")
	ErrAccountNotFound      = errors.New("the account is not found")
)

// AuthUseCase is an interface to provide features related to account authentication.
type AuthUseCase interface {
	// SignUp creates a new account with the given username and password.
	// This function expects that the values passed to it are validated.
	// When the username is already used, it returns an ErrUsernameConflict error.
	SignUp(ctx context.Context, username string, password string) (*Account, error)

	// SignIn tries to authenticate an account with the given username and password.
	// This function expects that the values passed to it are validated.
	// When the username or password is not correct, it returns an ErrAuthenticationFailed error.
	// ErrAuthenticationFailed is returned when the username is not found or the password is incorrect.
	SignIn(ctx context.Context, username string, password string) (*Account, error)
}

// AuthRepository is an repository interface used by AuthUseCase.
type AuthRepository interface {
	// IsUsernameUsed checks if the given username is already used.
	IsUsernameUsed(ctx context.Context, username string) (bool, error)

	// Save saves a new account to the repository.
	Save(ctx context.Context, account *Account) error

	// FindByUsername finds an account by the given username.
	// When the account is not found, it returns an ErrAccountNotFound error.
	FindByUsername(ctx context.Context, username string) (*Account, error)
}

// AuthService must be implemented by AuthUseCase.
var _ AuthUseCase = (*AuthService)(nil)

// AuthService is a implementation of AuthUseCase.
type AuthService struct {
	authRepository     AuthRepository
	encryptService     encrypt.EncryptService
	identifierProvider identifier.Provider
}

// NewAuthService creates a new AuthService with the given repository.
func NewAuthService(repo AuthRepository, enc encrypt.EncryptService, prov identifier.Provider) *AuthService {
	return &AuthService{
		authRepository:     repo,
		encryptService:     enc,
		identifierProvider: prov,
	}
}

func (s *AuthService) SignUp(ctx context.Context, username string, password string) (*Account, error) {
	if exist, err := s.authRepository.IsUsernameUsed(ctx, username); err != nil {
		return nil, oops.
			With("username", username).
			Errorf("failed to check if the username is used: %v", err)
	} else if exist {
		return nil, ErrUsernameConflict
	}

	hashed, err := s.encryptService.Encrypt([]byte(password))
	if err != nil {
		return nil, oops.
			Errorf("failed to encrypt the password: %v", err)
	}

	account := NewAccount(s.identifierProvider.Generate(), username, string(hashed))

	if err := s.authRepository.Save(ctx, account); err != nil {
		return nil, oops.
			With("account", account).
			Errorf("failed to save the account: %v", err)
	}
	return account, nil
}

func (s *AuthService) SignIn(ctx context.Context, username string, password string) (*Account, error) {
	account, err := s.authRepository.FindByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, ErrAccountNotFound) {
			return nil, err
		}
		return nil, oops.
			With("username", username).
			Errorf("failed to find the account by the username: %v", err)
	}

	if err := s.encryptService.Compare([]byte(account.PasswordHash), []byte(password)); err != nil {
		return nil, ErrAuthenticationFailed
	}

	return account, nil
}
