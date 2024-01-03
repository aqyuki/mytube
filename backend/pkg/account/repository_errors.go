package account

import "fmt"

// ErrInternalDatabase is an error type for internal database error
type ErrInternalDatabase struct {
	Err error
}

func (e *ErrInternalDatabase) Error() string {
	return fmt.Sprintf("internal database error: %s", e.Err.Error())
}

// NewInternalDatabaseErr creates a new internal database error
func NewInternalDatabaseErr(err error) *ErrInternalDatabase {
	return &ErrInternalDatabase{Err: err}
}

// IsInternalDatabaseErr checks if an error is an internal database error
func IsInternalDatabaseErr(err error) bool {
	_, ok := err.(*ErrInternalDatabase)
	return ok
}

// ErrAccountNotFound is an error type for account not found error
type ErrAccountNotFound struct {
	Username string
}

func (e *ErrAccountNotFound) Error() string {
	return fmt.Sprintf("%s is not found", e.Username)
}

// NewAccountNotFoundErr creates a new account not found error
func NewAccountNotFoundErr(username string) *ErrAccountNotFound {
	return &ErrAccountNotFound{Username: username}
}

// IsAccountNotFoundErr checks if an error is an account not found error
func IsAccountNotFoundErr(err error) bool {
	_, ok := err.(*ErrAccountNotFound)
	return ok
}

type ErrAccountAlreadyExists struct {
	Username string
}

func (e *ErrAccountAlreadyExists) Error() string {
	return fmt.Sprintf("%s already exists", e.Username)
}

// NewAccountAlreadyExistsErr creates a new account already exists error
func NewAccountAlreadyExistsErr(username string) *ErrAccountAlreadyExists {
	return &ErrAccountAlreadyExists{Username: username}
}

// IsAccountAlreadyExistsErr checks if an error is an account already exists error
func IsAccountAlreadyExistsErr(err error) bool {
	_, ok := err.(*ErrAccountAlreadyExists)
	return ok
}
