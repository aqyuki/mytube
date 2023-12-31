package account

import "context"

type AccountRepository interface {
	// FetchByUserName fetches an account by its username
	FetchByUserName(ctx context.Context, userName string) (*Account, error)
	// Register registers a new account to the repository
	Register(ctx context.Context, account *Account) error
	// Update updates an account in the repository
	Update(ctx context.Context, account *Account) error
	// Delete deletes an account from the repository
	Delete(ctx context.Context, account *Account) error
}
