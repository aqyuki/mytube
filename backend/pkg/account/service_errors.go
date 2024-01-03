package account

import "fmt"

type ErrInternalAccountService struct {
	Err error
}

func (e *ErrInternalAccountService) Error() string {
	return fmt.Sprintf("internal account service error: %s", e.Err.Error())
}

// NewInternalAccountServiceErr creates a new internal account service error
func NewInternalAccountServiceErr(err error) *ErrInternalAccountService {
	return &ErrInternalAccountService{Err: err}
}

// IsInternalAccountServiceErr checks if an error is an internal account service error
func IsInternalAccountServiceErr(err error) bool {
	_, ok := err.(*ErrInternalAccountService)
	return ok
}

type ErrAccountOAuth struct{}

func (e *ErrAccountOAuth) Error() string {
	return "account oauth error"
}

// NewAccountOAuthErr creates a new account oauth error
func NewAccountOAuthErr() *ErrAccountOAuth {
	return &ErrAccountOAuth{}
}

func IsErrAccountOAuth(err error) bool {
	_, ok := err.(*ErrAccountOAuth)
	return ok
}
