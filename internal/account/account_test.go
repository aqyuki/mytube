package account

import "testing"

func TestNewAccount(t *testing.T) {
	t.Parallel()

	id := "id"
	username := "username"
	passwordHash := "passwordHash"

	account := NewAccount(id, username, passwordHash)

	if account == nil {
		t.Fatal("NewAccount must be return an account but received nil")
	}
	if account.ID != id {
		t.Errorf("ID must be %q but got %q", id, account.ID)
	}
	if account.Username != username {
		t.Errorf("Username must be %q but got %q", username, account.Username)
	}
	if account.PasswordHash != passwordHash {
		t.Errorf("PasswordHash must be %q but got %q", passwordHash, account.PasswordHash)
	}
}
