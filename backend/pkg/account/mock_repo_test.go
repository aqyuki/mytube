package account

import (
	"context"
	"errors"
	"sync"
)

var _ Repository = (*MockAccountRepository)(nil)

type MockAccountRepository struct {
	m       map[string]Account
	mu      sync.RWMutex
	ErrMode bool
}

func (m *MockAccountRepository) FetchByUsername(ctx context.Context, userName string) (*Account, error) {
	if m.ErrMode {
		return nil, errors.New("error mode")
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	if account, ok := m.m[userName]; ok {
		return &account, nil
	}
	return nil, errors.New("account not found")
}

func (m *MockAccountRepository) Register(ctx context.Context, account *Account) error {
	if m.ErrMode {
		return errors.New("error mode")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.m[account.Username] = *account
	return nil
}

func (m *MockAccountRepository) Update(ctx context.Context, account *Account) error {
	if m.ErrMode {
		return errors.New("error mode")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.m[account.Username] = *account
	return nil
}

func (m *MockAccountRepository) Delete(ctx context.Context, account *Account) error {
	if m.ErrMode {
		return errors.New("error mode")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.m, account.Username)
	return nil
}

func NewMockAccountRepository() *MockAccountRepository {
	return &MockAccountRepository{
		m: make(map[string]Account),
	}
}
