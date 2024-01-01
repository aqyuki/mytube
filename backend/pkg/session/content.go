package session

import "time"

type Content struct {
	Username  string
	OAuth     bool
	CreatedAt int64
}

func (c *Content) ToMap() map[any]any {
	return map[any]any{
		"username":   c.Username,
		"oauth":      c.OAuth,
		"created_at": c.CreatedAt,
	}
}

func (c *Content) FromMap(m map[any]any) {
	c.Username = m["username"].(string)
	c.OAuth = m["oauth"].(bool)
	c.CreatedAt = m["created_at"].(int64)
}

func NewContent(username string, oauth bool, createdAt time.Time) *Content {
	return &Content{
		Username:  username,
		OAuth:     oauth,
		CreatedAt: createdAt.Unix(),
	}
}
