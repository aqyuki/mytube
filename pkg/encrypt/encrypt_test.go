package encrypt

import (
	"strings"
	"testing"
)

func TestNewBcryptEncryptService(t *testing.T) {
	t.Parallel()

	actual := NewBcryptEncryptService()
	if actual == nil {
		t.Error("NewBcryptEncryptService must not be nil")
	}
}

func TestBcryptEncryptService_Encrypt(t *testing.T) {
	t.Parallel()

	t.Run("should return encrypted byte slice", func(t *testing.T) {
		t.Parallel()

		s := NewBcryptEncryptService()
		actual, err := s.Encrypt([]byte("test"))
		if err != nil {
			t.Errorf("error should be nil, but got: %v", err)
		}
		if actual == nil {
			t.Error("actual should not be nil")
		}
	})

	t.Run("should return error when the longer password is given", func(t *testing.T) {
		t.Parallel()

		s := NewBcryptEncryptService()
		actual, err := s.Encrypt([]byte(strings.Repeat("a", 73)))
		if err == nil {
			t.Error("error should not be nil")
		}
		if actual != nil {
			t.Error("actual should be nil")
		}
	})
}

func TestBcryptEncryptService_Compare(t *testing.T) {
	t.Parallel()

	t.Run("should return nil when the hashed byte slice is the same as the target byte slice", func(t *testing.T) {
		t.Parallel()

		s := NewBcryptEncryptService()
		hashed, err := s.Encrypt([]byte("test"))
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}

		if err := s.Compare(hashed, []byte("test")); err != nil {
			t.Errorf("error should be nil, but got: %v", err)
		}
	})

	t.Run("should return ErrMissMatch when the hashed byte slice is different from the target byte slice", func(t *testing.T) {
		t.Parallel()

		s := NewBcryptEncryptService()
		hashed, err := s.Encrypt([]byte("test"))
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}

		if err := s.Compare(hashed, []byte("invalid")); err != ErrMissMatch {
			t.Errorf("error should be ErrMissMatch, but got: %v", err)
		}
	})
}
