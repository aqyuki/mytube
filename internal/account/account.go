package account

// Account is a struct that represents an account.
type Account struct {
	// ID is a string of characters that uniquely identifies an account.
	// It is a value used internally by the system and is not transmitted externally.
	// To uniquely identify an account outside the system, use the username.
	ID string `json:"id"`

	// Username is a string of characters that uniquely identifies an account.
	// A user name is a string that uniquely identifies an account from outside
	// the system, and user names, rather than IDs, should be used outside the system whenever possible.
	Username string `json:"username"`

	// PasswordHash is a string of characters that represents the hash of the account's password.
	PasswordHash string `json:"-"`
}

// NewAccount creates a new account model with the given ID, username, and password hash.
func NewAccount(id, username, passwordHash string) *Account {
	return &Account{
		ID:           id,
		Username:     username,
		PasswordHash: passwordHash,
	}
}
