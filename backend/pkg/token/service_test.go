package token

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTokenVerifier(t *testing.T) {
	t.Parallel()
	actual := NewTokenVerifier("secret", NewMockRepository())
	assert.NotNil(t, actual, "NewTokenVerifier should return a non-nil value")
}

func TestTokenVerifier_Generate(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		repository := NewMockRepository()
		verifier := NewTokenVerifier("secret", repository)

		token, err := verifier.Generate(context.Background(), "username")
		assert.NoError(t, err, "Generate should not return an error")
		assert.NotEmpty(t, token, "Generate should return a non-empty token")
	})
}
