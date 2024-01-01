package token

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/xid"
)

type TokenClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

type Service interface {
	// Generate generates a new JWT token for the given username
	Generate(ctx context.Context, username string) (string, error)
	// Verify verifies the given JWT token
	Verify(ctx context.Context, token, username string) error
}

type TokenVerifier struct {
	secret     []byte
	repository Repository
}

func NewTokenVerifier(secret string, repository Repository) *TokenVerifier {
	return &TokenVerifier{
		secret:     []byte(secret),
		repository: repository,
	}
}

func (t *TokenVerifier) Generate(ctx context.Context, username string) (string, error) {
	tokenID := xid.New().String()

	claims := &TokenClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	if err := t.repository.Register(ctx, tokenID, username); err != nil {
		return "", fmt.Errorf("failed registers tokenID: %w", err)
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(t.secret)
}

func (t *TokenVerifier) Verify(ctx context.Context, claims any, username string) error {
	c, ok := claims.(*TokenClaims)
	if !ok {
		return fmt.Errorf("invalid claims type")
	}

	pair, err := t.repository.FindByTokenID(ctx, c.ID)
	if err != nil {
		return fmt.Errorf("failed to find tokenID: %w", err)
	}

	if username != pair {
		return fmt.Errorf("not matching username")
	}
	return nil
}
