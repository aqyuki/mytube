package account

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type Account struct {
	bun.BaseModel `bun:"table:accounts"`
	ID            string       `bun:"id,pk"`
	Username      string       `bun:"username,notnull,unique"`
	PasswordHash  string       `bun:"password_hash,notnull"` // PasswordHash is the hash of the password
	CreatedAt     time.Time    `bun:"created_at"`
	UpdatedAt     time.Time    `bun:"updated_at"`
	DeletedAt     sql.NullTime `bun:"deleted_at,soft_delete,nullzero"`
}

// EqualPassword compares the password with the hash
func (a *Account) EqualPassword(password string) bool {
	return comparePassword(a.PasswordHash, password) == nil
}

// UpdateUsername updates the user name
func (a *Account) UpdateUsername(userName string) {
	a.Username = userName
	a.UpdatedAt = time.Now()
}

// UpdatePassword updates the password
func (a *Account) UpdatePassword(password string) error {
	hash, err := encryptPassword(password)
	if err != nil {
		return err
	}
	a.PasswordHash = hash
	a.UpdatedAt = time.Now()
	return nil
}

// NewAccount creates a new account
func NewAccount(id, user, password string, createdAt, updatedAt time.Time) (*Account, error) {
	hashedPassword, err := encryptPassword(password)
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:           id,
		Username:     user,
		PasswordHash: hashedPassword,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		DeletedAt:    sql.NullTime{},
	}, nil
}
