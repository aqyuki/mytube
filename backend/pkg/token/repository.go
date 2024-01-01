package token

import "context"

// Repository manages the tokenID and username pair
type Repository interface {
	// Register registers the tokenID and username pair
	Register(ctx context.Context, tokenID, username string) error
	// FindByTokenID finds the username by the given tokenID
	FindByTokenID(ctx context.Context, tokenID string) (string, error)
}
