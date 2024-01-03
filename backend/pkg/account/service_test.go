package account

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Register(t *testing.T) {
	t.Parallel()

	t.Run("success to register account", func(t *testing.T) {
		t.Parallel()

		repo := NewMockAccountRepository()
		repo.ErrMode = false

		service := NewAccountService(repo)
		ctx := context.Background()

		user, err := service.Register(ctx, "test", "test")
		assert.NoError(t, err, "should not return error")
		if assert.NotNil(t, user, "should return user") {
			assert.Equal(t, "test", user.Username, "should return correct username")
			user.EqualPassword("test")
		} else {
			t.Fatal("should return user")
		}
	})

	t.Run("failed to register account - already exists", func(t *testing.T) {
		t.Parallel()

		repo := NewMockAccountRepository()
		repo.ErrMode = false

		service := NewAccountService(repo)
		ctx := context.Background()

		user, err := service.Register(ctx, "test", "test")
		assert.NoError(t, err, "should not return error")
		if assert.NotNil(t, user, "should return user") {
			assert.Equal(t, "test", user.Username, "should return correct username")
			user.EqualPassword("test")
		} else {
			t.Fatal("should return user")
		}

		user, err = service.Register(ctx, "test", "test")
		assert.True(t, IsAccountAlreadyExistsErr(err), "should return account already exists error")
		assert.Nil(t, user, "should not return user")
	})

	t.Run("failed to register account - internal repository error", func(t *testing.T) {
		t.Parallel()

		repo := NewMockAccountRepository()
		repo.ErrMode = true

		service := NewAccountService(repo)
		ctx := context.Background()

		user, err := service.Register(ctx, "test", "test")
		assert.Error(t, err, "should return error")
		assert.Nil(t, user, "should not return user")
	})
}

func TestService_Delete(t *testing.T) {
	t.Parallel()

	t.Run("success to delete account", func(t *testing.T) {
		t.Parallel()

		repo := NewMockAccountRepository()
		repo.ErrMode = false

		service := NewAccountService(repo)
		ctx := context.Background()

		user, err := service.Register(ctx, "test", "test")
		assert.NoError(t, err, "should not return error")

		err = service.Delete(ctx, user.Username)
		assert.NoError(t, err, "should not return error")
	})

	t.Run("failed to delete account - not found account", func(t *testing.T) {
		t.Parallel()

		repo := NewMockAccountRepository()
		repo.ErrMode = false

		service := NewAccountService(repo)

		err := service.Delete(context.Background(), "test")
		assert.True(t, IsAccountNotFoundErr(err), "should return account not found error")
	})

	t.Run("failed to delete account - internal repository error", func(t *testing.T) {
		t.Parallel()

		repo := NewMockAccountRepository()
		repo.ErrMode = false

		service := NewAccountService(repo)
		ctx := context.Background()

		user, err := service.Register(ctx, "test", "test")
		assert.NoError(t, err, "should not return error")

		repo.ErrMode = true
		err = service.Delete(ctx, user.Username)
		assert.True(t, IsInternalAccountServiceErr(err), "should return internal account service error")
	})
}

func TestService_Login(t *testing.T) {
	t.Parallel()

	t.Run("success to login", func(t *testing.T) {
		t.Parallel()

		repo := NewMockAccountRepository()
		repo.ErrMode = false

		service := NewAccountService(repo)
		ctx := context.Background()

		user, err := service.Register(ctx, "test", "test")
		assert.NoError(t, err, "should not return error")
		if assert.NotNil(t, user, "should return user") {
			assert.Equal(t, "test", user.Username, "should return correct username")
			user.EqualPassword("test")
		} else {
			t.Fatal("should return user")
		}
		user, err = service.Login(ctx, "test", "test")
		assert.NoError(t, err, "should not return error")
		if assert.NotNil(t, user, "should return user") {
			assert.Equal(t, "test", user.Username, "should return correct username")
			user.EqualPassword("test")
		} else {
			t.Fatal("should return user")
		}
	})

	t.Run("failed to login - internal repository error", func(t *testing.T) {
		t.Parallel()

		repo := NewMockAccountRepository()
		repo.ErrMode = false

		service := NewAccountService(repo)
		ctx := context.Background()

		user, err := service.Register(ctx, "test", "test")
		assert.NoError(t, err, "should not return error")
		if assert.NotNil(t, user, "should return user") {
			assert.Equal(t, "test", user.Username, "should return correct username")
			user.EqualPassword("test")
		} else {
			t.Fatal("should return user")
		}

		repo.ErrMode = true
		user, err = service.Login(ctx, "test", "test")
		assert.True(t, IsInternalAccountServiceErr(err), "should return internal account service error")
		assert.Nil(t, user, "should not return user")
	})

	t.Run("failed to oauth - not found account", func(t *testing.T) {
		t.Parallel()

		repo := NewMockAccountRepository()
		repo.ErrMode = false

		service := NewAccountService(repo)

		user, err := service.Login(context.Background(), "test", "test")
		assert.True(t, IsErrAccountOAuth(err), "should return account oauth error")
		assert.Nil(t, user, "should not return user")
	})

	t.Run("failed to oauth - not match password", func(t *testing.T) {
		t.Parallel()

		repo := NewMockAccountRepository()
		repo.ErrMode = false

		service := NewAccountService(repo)
		ctx := context.Background()

		user, err := service.Register(ctx, "test", "test")
		assert.NoError(t, err, "should not return error")
		if assert.NotNil(t, user, "should return user") {
			assert.Equal(t, "test", user.Username, "should return correct username")
			user.EqualPassword("test")
		} else {
			t.Fatal("should return user")
		}

		user, err = service.Login(ctx, "test", "test2")
		assert.True(t, IsErrAccountOAuth(err), "should return account oauth error")
		assert.Nil(t, user, "should not return user")
	})
}
