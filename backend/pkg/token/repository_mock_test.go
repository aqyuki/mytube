package token

import (
	"context"
	"errors"
	"sync"
)

type MockRepository struct {
	m  map[string]string
	mu sync.RWMutex
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		m: make(map[string]string),
	}
}

func (r *MockRepository) Register(_ context.Context, tokenID, username string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.m[tokenID] = username
	return nil
}

func (r *MockRepository) FindByTokenID(_ context.Context, tokenID string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	username, ok := r.m[tokenID]
	if !ok {
		return "", errors.New("not found")
	}
	return username, nil
}
