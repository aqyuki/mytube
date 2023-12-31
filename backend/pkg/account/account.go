package account

import (
	"time"
)

type Account struct {
	ID           string
	UserName     string
	PasswordHash string // PasswordHash is the hash of the password
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// EqualPassword compares the password with the hash
func (a *Account) EqualPassword(password string) bool {
	return comparePassword(a.PasswordHash, password) == nil
}

// NewAccount creates a new account
func NewAccount(id, user, password string, createdAt, updatedAt time.Time) (*Account, error) {
	hashedPassword, err := encryptPassword(password)
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:           id,
		UserName:     user,
		PasswordHash: hashedPassword,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}, nil
}
