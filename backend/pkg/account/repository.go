package account

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type Repository interface {
	// FetchByUsername fetches an account by its username
	FetchByUsername(ctx context.Context, userName string) (*Account, error)
	// Register registers a new account to the repository
	Register(ctx context.Context, account *Account) error
	// Update updates an account in the repository
	Update(ctx context.Context, account *Account) error
	// Delete deletes an account from the repository
	Delete(ctx context.Context, account *Account) error
}

type AccountRepository struct {
	db *bun.DB
}

// NewAccountRepository creates a new account repository
func NewAccountRepository(sqlDB *sql.DB) *AccountRepository {
	db := bun.NewDB(sqlDB, pgdialect.New())
	return &AccountRepository{
		db: db,
	}
}

// FetchByUsername fetches an account by its username
func (r *AccountRepository) FetchByUsername(ctx context.Context, userName string) (*Account, error) {
	var user Account
	if err := r.db.NewSelect().Model(&user).Where("username = ?", userName).Scan(ctx); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AccountRepository) Register(ctx context.Context, account *Account) error {
	if _, err := r.db.NewInsert().Model(account).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (r *AccountRepository) Update(ctx context.Context, account *Account) error {
	if _, err := r.db.NewUpdate().Model(account).Where("id = ?", account.ID).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (r *AccountRepository) Delete(ctx context.Context, account *Account) error {
	if _, err := r.db.NewDelete().Model(account).Where("id = ?", account.ID).Exec(ctx); err != nil {
		return err
	}
	return nil
}
