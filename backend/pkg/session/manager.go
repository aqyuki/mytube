package session

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

const (
	sessionKey     = "mt-session"
	redisKeyPrefix = "mt-session-"
)

var (
	ErrNotFound = errors.New("session not found")
)

type Manager struct {
	conn *redis.Client
}

func (m *Manager) GetContent(c echo.Context) (*Content, error) {
	cookie, err := c.Cookie(sessionKey)
	if errors.Is(err, http.ErrNoCookie) {
		return &Content{
			OAuth: false,
		}, nil
	} else if err != nil {
		return nil, err
	}

	value, err := m.conn.Get(c.Request().Context(), cookie.Value).Result()
	if err != nil {
		return nil, ErrNotFound
	}

	content, err := convertToContent([]byte(value))
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (m *Manager) SaveContent(c echo.Context, content *Content) error {
	cookie, err := c.Cookie(sessionKey)
	if err != nil && !errors.Is(err, http.ErrNoCookie) {
		return err
	}

	var key string

	if errors.Is(err, http.ErrNoCookie) {
		genKey, err := generateSessionID()
		if err != nil {
			return err
		}
		key = genKey
	} else {
		key = cookie.Value
	}

	value, err := convertToJSON(content)
	if err != nil {
		return err
	}

	if err := m.conn.Set(c.Request().Context(), key, value, 0).Err(); err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     sessionKey,
		Path:     "/",
		Value:    key,
		HttpOnly: true,
		MaxAge:   int(30 * 24 * 60 * 60), // 1 month
	})
	return nil

}

func (m *Manager) DeleteContent(c echo.Context) error {
	cookie, err := c.Cookie(sessionKey)
	if err != nil && !errors.Is(err, http.ErrNoCookie) {
		return err
	} else if errors.Is(err, http.ErrNoCookie) {
		return nil
	}

	if err := m.conn.Del(c.Request().Context(), cookie.Value).Err(); err != nil {
		return err
	}

	return nil
}

func NewManager(client *redis.Client) *Manager {
	return &Manager{
		conn: client,
	}
}

func convertToJSON(content *Content) ([]byte, error) {
	return json.Marshal(content)
}

func convertToContent(data []byte) (*Content, error) {
	var content Content
	err := json.Unmarshal(data, &content)
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func generateSessionID() (string, error) {
	randomBytes := make([]byte, 128)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return redisKeyPrefix + base64.URLEncoding.EncodeToString(randomBytes)[:172], nil
}
