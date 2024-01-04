package collection

import (
	"fmt"
	"time"
)

// Video holds a information about a vide registered in the mytube.
type Video struct {
	ID             string
	VideoID        string
	Title          string
	Description    string
	ChannelName    string
	ChannelIconURL string
	ChannelURL     string
	Fab            bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time // nullable
}

// Favorite marks a video as favorite.
func (v *Video) Favorite() {
	v.Fab = true
}

// Unfavorite marks a video as unfavorite.
func (v *Video) Unfavorite() {
	v.Fab = false
}

// String returns a string representation of the video.
func (v *Video) String() string {
	var deletedAt string
	if v.DeletedAt == nil {
		deletedAt = "<empty>"
	} else {
		deletedAt = v.DeletedAt.Format(time.DateTime)
	}
	return fmt.Sprintf(
		"ID : %s\tVideoID : %s\tTitle : %s\tDescription : %s\tChannelName : %s\tChannelIconURL : %s\tChannelURL : %s\tFab : %t\tCreatedAt : %s\tUpdatedAt : %s\tDeletedAt : %s",
		v.ID, v.VideoID, v.Title, v.Description, v.ChannelName, v.ChannelIconURL, v.ChannelURL, v.Fab, v.CreatedAt.Format(time.DateTime), v.UpdatedAt.Format(time.DateTime), deletedAt,
	)
}

// NewVideo creates a new Video instance.
func NewVideo(id, videoID, title, description, channelName, channelIconURL, channelURL string, fab bool, createdAt, updatedAt time.Time) *Video {
	return &Video{
		ID:             id,
		VideoID:        videoID,
		Title:          title,
		Description:    description,
		ChannelName:    channelName,
		ChannelIconURL: channelIconURL,
		ChannelURL:     channelURL,
		Fab:            fab,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		DeletedAt:      nil,
	}
}
