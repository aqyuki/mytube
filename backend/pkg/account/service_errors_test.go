package account

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrInternalAccountService(t *testing.T) {
	t.Parallel()

	t.Run("Same Error", func(t *testing.T) {
		t.Parallel()

		actual := NewInternalAccountServiceErr(errors.New("test"))

		assert.Error(t, actual, "should be an error")
		assert.EqualError(t, actual, "internal account service error: test")
		assert.True(t, IsInternalAccountServiceErr(actual), "should be internal account service error")
	})

	t.Run("Different Error", func(t *testing.T) {
		t.Parallel()

		actual := NewInternalAccountServiceErr(errors.New("test"))

		assert.Error(t, actual, "should be an error")
		assert.EqualError(t, actual, "internal account service error: test")
		assert.False(t, IsInternalAccountServiceErr(errors.New("error")), "should not be internal account service error")
	})
}

func TestErrAccountOAuth(t *testing.T) {
	t.Parallel()

	t.Run("Same Error", func(t *testing.T) {
		t.Parallel()

		actual := NewAccountOAuthErr()

		assert.Error(t, actual, "should be an error")
		assert.EqualError(t, actual, "account oauth error")
		assert.True(t, IsErrAccountOAuth(actual), "should be account oauth error")
	})

	t.Run("Different Error", func(t *testing.T) {
		t.Parallel()

		actual := NewAccountOAuthErr()

		assert.Error(t, actual, "should be an error")
		assert.EqualError(t, actual, "account oauth error")
		assert.NotEqual(t, "test", actual.Error())
	})
}
