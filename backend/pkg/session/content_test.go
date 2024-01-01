package session

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewContent(t *testing.T) {
	t.Parallel()

	data := struct {
		username  string
		oauth     bool
		createdAt time.Time
	}{
		username:  "test",
		oauth:     true,
		createdAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	content := NewContent(data.username, data.oauth, data.createdAt)
	assert.EqualValues(t, data.username, content.Username)
	assert.EqualValues(t, data.oauth, content.OAuth)
	assert.EqualValues(t, data.createdAt.Unix(), content.CreatedAt)
}

func TestContent_ToMap(t *testing.T) {
	t.Parallel()

	data := struct {
		username  string
		oauth     bool
		createdAt time.Time
	}{
		username:  "test",
		oauth:     true,
		createdAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	content := NewContent(data.username, data.oauth, data.createdAt)
	m := content.ToMap()
	assert.EqualValues(t, data.username, m["username"])
	assert.EqualValues(t, data.oauth, m["oauth"])
	assert.EqualValues(t, data.createdAt.Unix(), m["created_at"])
}

func TestContent_FromMap(t *testing.T) {
	t.Parallel()

	data := struct {
		username  string
		oauth     bool
		createdAt time.Time
	}{
		username:  "test",
		oauth:     true,
		createdAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	content := NewContent(data.username, data.oauth, data.createdAt)
	m := content.ToMap()

	content = &Content{}
	content.FromMap(m)
	assert.EqualValues(t, data.username, content.Username)
	assert.EqualValues(t, data.oauth, content.OAuth)
	assert.EqualValues(t, data.createdAt.Unix(), content.CreatedAt)
}