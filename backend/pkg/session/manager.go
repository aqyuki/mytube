package session

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/rbcervilla/redisstore/v9"
)

const sessionKey = "mt-session"

type Manager struct {
	store *redisstore.RedisStore
}

func (m *Manager) GetContent(c echo.Context) (*Content, error) {
	session, err := m.store.Get(c.Request(), sessionKey)
	if err != nil {
		return nil, err
	}

	var content Content
	content.FromMap(session.Values)
	return &content, nil
}

func (m *Manager) SaveContent(c echo.Context, content *Content) error {
	session, err := m.store.Get(c.Request(), sessionKey)
	if err != nil {
		return err
	}
	if content == nil {
		return errors.New("failed to save session: nil content")
	}
	session.Values = content.ToMap()
	return session.Save(c.Request(), c.Response())
}

func NewManager(store *redisstore.RedisStore) *Manager {
	return &Manager{
		store: store,
	}
}
