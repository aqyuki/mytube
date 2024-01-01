package account

import (
	"time"
)

type Account struct {
	ID           string
	Username     string
	PasswordHash string // PasswordHash is the hash of the password
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
	}, nil
}
