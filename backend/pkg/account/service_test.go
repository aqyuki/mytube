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
			assert.Equal(t, "test", user.UserName, "should return correct username")
			user.EqualPassword("test")
		} else {
			t.Fatal("should return user")
		}
	})

	t.Run("failed to register account", func(t *testing.T) {
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

		err = service.Delete(ctx, user.UserName)
		assert.NoError(t, err, "should not return error")
	})

	t.Run("failed to delete account", func(t *testing.T) {
		t.Parallel()

		repo := NewMockAccountRepository()
		repo.ErrMode = false

		service := NewAccountService(repo)
		ctx := context.Background()

		user, err := service.Register(ctx, "test", "test")
		assert.NoError(t, err, "should not return error")

		repo.ErrMode = true
		err = service.Delete(ctx, user.UserName)
		assert.Error(t, err, "should return error")
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
			assert.Equal(t, "test", user.UserName, "should return correct username")
			user.EqualPassword("test")
		} else {
			t.Fatal("should return user")
		}
		user, err = service.Login(ctx, "test", "test")
		assert.NoError(t, err, "should not return error")
		if assert.NotNil(t, user, "should return user") {
			assert.Equal(t, "test", user.UserName, "should return correct username")
			user.EqualPassword("test")
		} else {
			t.Fatal("should return user")
		}
	})

	t.Run("failed to login", func(t *testing.T) {
		t.Parallel()

		repo := NewMockAccountRepository()
		repo.ErrMode = false

		service := NewAccountService(repo)
		ctx := context.Background()

		user, err := service.Register(ctx, "test", "test")
		assert.NoError(t, err, "should not return error")
		if assert.NotNil(t, user, "should return user") {
			assert.Equal(t, "test", user.UserName, "should return correct username")
			user.EqualPassword("test")
		} else {
			t.Fatal("should return user")
		}
		user, err = service.Login(ctx, "test", "test2")
		assert.Error(t, err, "should return error")
		assert.Nil(t, user, "should not return user")
	})
}
