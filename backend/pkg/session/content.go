package session

import "time"

type Content struct {
	Username  string `json:"username"`
	OAuth     bool   `json:"oauth"`
	CreatedAt int64  `json:"created_at"`
}

func NewContent(username string, oauth bool, createdAt time.Time) *Content {
	return &Content{
		Username:  username,
		OAuth:     oauth,
		CreatedAt: createdAt.Unix(),
	}
}
