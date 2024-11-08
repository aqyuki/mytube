package account

import (
	"context"
	"errors"
	"testing"

	"github.com/aqyuki/mytube/pkg/encrypt"
	"github.com/aqyuki/mytube/pkg/identifier"
)

var _ AuthRepository = (*MockAuthRepository)(nil)

type MockAuthRepository struct {
	IsUsernameUsedFn func(ctx context.Context, username string) (bool, error)
	SaveFn           func(ctx context.Context, account *Account) error
	FindByUsernameFn func(ctx context.Context, username string) (*Account, error)
}

func (m *MockAuthRepository) IsUsernameUsed(ctx context.Context, username string) (bool, error) {
	return m.IsUsernameUsedFn(ctx, username)
}

func (m *MockAuthRepository) Save(ctx context.Context, account *Account) error {
	return m.SaveFn(ctx, account)
}

func (m *MockAuthRepository) FindByUsername(ctx context.Context, username string) (*Account, error) {
	return m.FindByUsernameFn(ctx, username)
}

var _ encrypt.EncryptService = (*MockEncryptService)(nil)

type MockEncryptService struct {
	EncryptFn func(b []byte) ([]byte, error)
	CompareFn func(hashed, target []byte) error
}

func (m *MockEncryptService) Encrypt(b []byte) ([]byte, error) {
	return m.EncryptFn(b)
}

func (m *MockEncryptService) Compare(hashed, target []byte) error {
	return m.CompareFn(hashed, target)
}

var _ identifier.Provider = (*MockIdentifierProvider)(nil)

type MockIdentifierProvider struct {
	GenerateFn func() string
}

func (m *MockIdentifierProvider) Generate() string {
	return m.GenerateFn()
}

func TestNewAuthService(t *testing.T) {
	t.Parallel()

	repo := new(MockAuthRepository)
	enc := new(MockEncryptService)
	prov := new(MockIdentifierProvider)

	actual := NewAuthService(repo, enc, prov)

	if actual == nil {
		t.Error("NewAuthService must be returned a not nil value but received nil")
	}
}

func TestAuthService_SignUp(t *testing.T) {
	t.Parallel()

	t.Run("should return an account and nil error", func(t *testing.T) {
		t.Parallel()

		repo := new(MockAuthRepository)
		enc := new(MockEncryptService)
		prov := new(MockIdentifierProvider)

		repo.IsUsernameUsedFn = func(ctx context.Context, username string) (bool, error) {
			if username != "bar" {
				t.Errorf("IsUsernameUsed must be called with bar but called with %s", username)
			}

			return false, nil
		}
		repo.SaveFn = func(ctx context.Context, account *Account) error {
			if account == nil {
				t.Fatal("Save must be called with account but called with nil")
			}

			if account.ID != "foo" {
				t.Errorf("Save must be called with foo but called with %s", account.ID)
			}
			if account.Username != "bar" {
				t.Errorf("Save must be called with bar but called with %s", account.Username)
			}
			if account.PasswordHash != "hashed" {
				t.Errorf("Save must be called with hashed but called with	%s", account.PasswordHash)
			}

			return nil
		}
		enc.EncryptFn = func(b []byte) ([]byte, error) {
			if string(b) != "baz" {
				t.Fatalf("Encrypt must be called with baz but called with %s", string(b))
			}
			return []byte("hashed"), nil
		}
		prov.GenerateFn = func() string {
			return "foo"
		}

		srv := NewAuthService(repo, enc, prov)

		account, err := srv.SignUp(context.Background(), "bar", "baz")

		if err != nil {
			t.Fatalf("SignUp must return nil error but received %v", err)
		}
		if account == nil {
			t.Fatal("SignUp must return an account but received nil")
		}
		if account.ID != "foo" {
			t.Errorf("SignUp must return an account with foo but received %s", account.ID)
		}
		if account.Username != "bar" {
			t.Errorf("SignUp must return an account with bar but received %s", account.Username)
		}
		if account.PasswordHash != "hashed" {
			t.Errorf("SignUp must return an account with hashed but received %s", account.PasswordHash)
		}
	})

	t.Run("should return ErrUsernameConflict", func(t *testing.T) {
		t.Parallel()

		repo := new(MockAuthRepository)
		enc := new(MockEncryptService)
		prov := new(MockIdentifierProvider)

		repo.IsUsernameUsedFn = func(ctx context.Context, username string) (bool, error) {
			return true, nil
		}

		srv := NewAuthService(repo, enc, prov)

		_, err := srv.SignUp(context.Background(), "bar", "baz")

		if !errors.Is(err, ErrUsernameConflict) {
			t.Fatalf("SignUp must return ErrUsernameConflict but received %v", err)
		}
	})

	t.Run("should return an error when IsUsernameUsed returns an error", func(t *testing.T) {
		t.Parallel()

		repo := new(MockAuthRepository)
		enc := new(MockEncryptService)
		prov := new(MockIdentifierProvider)

		repo.IsUsernameUsedFn = func(ctx context.Context, username string) (bool, error) {
			return false, errors.New("an error")
		}

		srv := NewAuthService(repo, enc, prov)

		_, err := srv.SignUp(context.Background(), "bar", "baz")
		if err == nil {
			t.Fatal("SignUp must return an error but received nil")
		}
	})

	t.Run("should return an error when Encrypt returns an error", func(t *testing.T) {
		t.Parallel()

		repo := new(MockAuthRepository)
		enc := new(MockEncryptService)
		prov := new(MockIdentifierProvider)

		repo.IsUsernameUsedFn = func(ctx context.Context, username string) (bool, error) {
			return false, nil
		}
		enc.EncryptFn = func(b []byte) ([]byte, error) {
			return nil, errors.New("an error")
		}

		srv := NewAuthService(repo, enc, prov)

		_, err := srv.SignUp(context.Background(), "bar", "baz")
		if err == nil {
			t.Fatal("SignUp must return an error but received nil")
		}
	})

	t.Run("should return an error when Save returns an error", func(t *testing.T) {
		t.Parallel()

		repo := new(MockAuthRepository)
		enc := new(MockEncryptService)
		prov := new(MockIdentifierProvider)

		repo.IsUsernameUsedFn = func(ctx context.Context, username string) (bool, error) {
			return false, nil
		}
		repo.SaveFn = func(ctx context.Context, account *Account) error {
			return errors.New("an error")
		}
		enc.EncryptFn = func(b []byte) ([]byte, error) {
			return []byte("hashed"), nil
		}
		prov.GenerateFn = func() string {
			return "foo"
		}

		srv := NewAuthService(repo, enc, prov)

		_, err := srv.SignUp(context.Background(), "bar", "baz")
		if err == nil {
			t.Fatal("SignUp must return an error but received nil")
		}
	})
}

func TestAuthService_SignIn(t *testing.T) {
	t.Parallel()

	t.Run("should return an account and nil error", func(t *testing.T) {
		t.Parallel()

		repo := new(MockAuthRepository)
		enc := new(MockEncryptService)
		prov := new(MockIdentifierProvider)

		repo.FindByUsernameFn = func(ctx context.Context, username string) (*Account, error) {
			if username != "bar" {
				t.Errorf("FindByUsername must be called with bar but called with %s", username)
			}

			return &Account{
				ID:           "foo",
				Username:     "bar",
				PasswordHash: "hashed",
			}, nil
		}
		enc.CompareFn = func(hashed, target []byte) error {
			if string(hashed) != "hashed" {
				t.Fatalf("Compare must be called with hashed but called with %s", string(hashed))
			}
			if string(target) != "baz" {
				t.Fatalf("Compare must be called with baz but called with %s", string(target))
			}

			return nil
		}

		srv := NewAuthService(repo, enc, prov)
		account, err := srv.SignIn(context.Background(), "bar", "baz")

		if err != nil {
			t.Fatalf("SignIn must return nil error but received %v", err)
		}
		if account == nil {
			t.Fatal("SignIn must return an account but received nil")
		}
		if account.ID != "foo" {
			t.Errorf("SignIn must return an account with foo but received %s", account.ID)
		}
		if account.Username != "bar" {
			t.Errorf("SignIn must return an account with bar but received %s", account.Username)
		}
		if account.PasswordHash != "hashed" {
			t.Errorf("SignIn must return an account with hashed but received %s", account.PasswordHash)
		}
	})

	t.Run("should return ErrAccountNotFound", func(t *testing.T) {
		t.Parallel()

		repo := new(MockAuthRepository)
		enc := new(MockEncryptService)
		prov := new(MockIdentifierProvider)

		repo.FindByUsernameFn = func(ctx context.Context, username string) (*Account, error) {
			return nil, ErrAccountNotFound
		}

		srv := NewAuthService(repo, enc, prov)
		_, err := srv.SignIn(context.Background(), "bar", "baz")

		if !errors.Is(err, ErrAccountNotFound) {
			t.Fatalf("SignIn must return ErrAccountNotFound but received %v", err)
		}
	})

	t.Run("should return an error when FindByUsername returns an error", func(t *testing.T) {
		t.Parallel()

		repo := new(MockAuthRepository)
		enc := new(MockEncryptService)
		prov := new(MockIdentifierProvider)

		repo.FindByUsernameFn = func(ctx context.Context, username string) (*Account, error) {
			return nil, errors.New("an error")
		}

		srv := NewAuthService(repo, enc, prov)
		_, err := srv.SignIn(context.Background(), "bar", "baz")

		if err == nil {
			t.Fatal("SignIn must return an error but received nil")
		}
	})

	t.Run("should return an ErrAuthenticationFailed when Compare returns an error", func(t *testing.T) {
		t.Parallel()

		repo := new(MockAuthRepository)
		enc := new(MockEncryptService)
		prov := new(MockIdentifierProvider)

		repo.FindByUsernameFn = func(ctx context.Context, username string) (*Account, error) {
			return &Account{
				ID:           "foo",
				Username:     "bar",
				PasswordHash: "hashed",
			}, nil
		}
		enc.CompareFn = func(hashed, target []byte) error {
			return errors.New("an error")
		}

		srv := NewAuthService(repo, enc, prov)
		_, err := srv.SignIn(context.Background(), "bar", "baz")

		if !errors.Is(err, ErrAuthenticationFailed) {
			t.Fatalf("SignIn must return ErrAuthenticationFailed but received %v", err)
		}
	})
}
