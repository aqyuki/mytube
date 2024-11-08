package encrypt

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// EncryptService provides encryption-related processing.
type EncryptService interface {
	// Encrypt encrypts the given byte slice and returns the encrypted byte slice.
	Encrypt(b []byte) ([]byte, error)
	// Compare compares the given hashed byte slice with the target byte slice.
	// When the hashed byte slice is the same as the target byte slice, it returns nil.
	// When the hashed byte slice is different from the target byte slice, it returns an ErrMissMatch error.
	Compare(hashed, target []byte) error
}

// ErrMissMatch is an error that occurs when the hashed byte slice is different from the target byte slice.
var ErrMissMatch = errors.New("hashed byte slice is different from the target byte slice")

// BcryptEncryptService must be implemented by EncryptService.
var _ EncryptService = (*BcryptEncryptService)(nil)

// BcryptEncryptService is an implementation of EncryptService using bcrypt.
type BcryptEncryptService struct{}

// NewBcryptEncryptService creates a new BcryptEncryptService.
func NewBcryptEncryptService() *BcryptEncryptService { return new(BcryptEncryptService) }

// Encrypt encrypts the given byte slice using bcrypt.
func (s *BcryptEncryptService) Encrypt(b []byte) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword(b, 10)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt: %w", err)
	}
	return hashed, nil
}

// Compare compares the given hashed byte slice with the target byte slice using bcrypt.
func (s *BcryptEncryptService) Compare(hashed, target []byte) error {
	if err := bcrypt.CompareHashAndPassword(hashed, target); err != nil {
		return ErrMissMatch
	}
	return nil
}
