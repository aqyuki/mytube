package account

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	t.Parallel()

	user, err := NewAccount("id", "user", "password", time.Now(), time.Now())
	assert.NoError(t, err, "NewAccount should not return an error")
	if assert.NotNil(t, user, "NewAccount should return a user") {
		assert.EqualValues(t, "id", user.ID, "NewAccount should return the correct id")
		assert.EqualValues(t, "user", user.UserName, "NewAccount should return the correct username")
		assert.NotEmpty(t, user.PasswordHash, "NewAccount should return a password hash")
		assert.NotEmpty(t, user.CreatedAt, "NewAccount should return a created at timestamp")
		assert.NotEmpty(t, user.UpdatedAt, "NewAccount should return an updated at timestamp")
	}
}

func TestAccount_EqualPassword(t *testing.T) {
	t.Parallel()

	t.Run("should return true if the password is correct", func(t *testing.T) {
		user, err := NewAccount("id", "user", "password", time.Now(), time.Now())
		assert.NoError(t, err, "NewAccount should not return an error")
		assert.True(t, user.EqualPassword("password"))
	})

	t.Run("should return false if the password is incorrect", func(t *testing.T) {
		user, err := NewAccount("id", "user", "password", time.Now(), time.Now())
		assert.NoError(t, err, "NewAccount should not return an error")
		assert.False(t, user.EqualPassword("wrong password"))
	})
}

func TestAccount_UpdateUserName(t *testing.T) {
	t.Parallel()

	t.Run("should update the username", func(t *testing.T) {
		t.Parallel()

		user, err := NewAccount("id", "user", "password", time.Now(), time.Now())
		assert.NoError(t, err, "NewAccount should not return an error")
		user.UpdateUserName("new user")
		assert.EqualValues(t, "new user", user.UserName, "UpdateUserName should update the username")
	})
}

func TestAccount_UpdatePassword(t *testing.T) {
	t.Parallel()

	user, err := NewAccount("id", "user", "password", time.Now(), time.Now())
	assert.NoError(t, err, "NewAccount should not return an error")

	err = user.UpdatePassword("new password")
	assert.NoError(t, err, "UpdatePassword should not return an error")
	assert.True(t, user.EqualPassword("new password"), "UpdatePassword should update the password")
}
