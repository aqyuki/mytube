package account

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrInternalDatabase(t *testing.T) {
	t.Parallel()

	t.Run("Same Error", func(t *testing.T) {
		t.Parallel()

		err := errors.New("test")
		actual := NewInternalDatabaseErr(err)

		assert.Error(t, actual, "should be an error")
		assert.Equal(t, "internal database error: test", actual.Error())
		assert.True(t, IsInternalDatabaseErr(actual), "should be an internal database error")
	})

	t.Run("Different Error", func(t *testing.T) {
		t.Parallel()

		err := errors.New("test")
		actual := NewInternalDatabaseErr(err)

		assert.Error(t, actual, "should be an error")
		assert.Equal(t, "internal database error: test", actual.Error())
		assert.NotEqual(t, "test", actual.Error())
	})
}

func TestErrAccountNotFound(t *testing.T) {
	t.Parallel()

	t.Run("Same Error", func(t *testing.T) {
		t.Parallel()

		actual := NewAccountNotFoundErr("username")

		assert.Error(t, actual, "should be an error")
		assert.Equal(t, "username is not found", actual.Error())
		assert.True(t, IsAccountNotFoundErr(actual), "should be an account not found error")
	})

	t.Run("Different Error", func(t *testing.T) {
		t.Parallel()

		actual := NewAccountNotFoundErr("username")

		assert.Error(t, actual, "should be an error")
		assert.Equal(t, "username is not found", actual.Error())
		assert.NotEqual(t, "username", actual.Error())
	})
}

func TestErrAccountAlreadyExists(t *testing.T) {
	t.Parallel()

	t.Run("Same Error", func(t *testing.T) {
		t.Parallel()

		actual := NewAccountAlreadyExistsErr("username")

		assert.Error(t, actual, "should be an error")
		assert.Equal(t, "username already exists", actual.Error())
		assert.True(t, IsAccountAlreadyExistsErr(actual), "should be an account already exists error")
	})

	t.Run("Different Error", func(t *testing.T) {
		t.Parallel()

		actual := NewAccountAlreadyExistsErr("username")

		assert.Error(t, actual, "should be an error")
		assert.Equal(t, "username already exists", actual.Error())
		assert.NotEqual(t, "username", actual.Error())
	})
}
