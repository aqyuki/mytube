package collection

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func PointerTime(t *testing.T) *time.Time {
	t.Helper()
	date := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	return &date
}

func TestNewVideo(t *testing.T) {
	t.Parallel()

	type args struct {
		id             string
		videoID        string
		title          string
		description    string
		channelName    string
		channelIconURL string
		channelURL     string
		fab            bool
		createdAt      time.Time
		updatedAt      time.Time
	}

	tests := []struct {
		name string
		args args
		want *Video
	}{
		{
			name: "success",
			args: args{
				id:             "id",
				videoID:        "videoID",
				title:          "title",
				description:    "description",
				channelName:    "channelName",
				channelIconURL: "channelIconURL",
				channelURL:     "channelURL",
				fab:            true,
				createdAt:      time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				updatedAt:      time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			actual := NewVideo(tt.args.id, tt.args.videoID, tt.args.title, tt.args.description, tt.args.channelName, tt.args.channelIconURL, tt.args.channelURL, tt.args.fab, tt.args.createdAt, tt.args.updatedAt)
			if !assert.NotNil(t, actual, "NewVideo should not return nil, but received nil") {
				t.Fatal()
			}
			assert.EqualValues(t, tt.args.id, actual.ID)
			assert.EqualValues(t, tt.args.videoID, actual.VideoID)
			assert.EqualValues(t, tt.args.title, actual.Title)
			assert.EqualValues(t, tt.args.description, actual.Description)
			assert.EqualValues(t, tt.args.channelName, actual.ChannelName)
			assert.EqualValues(t, tt.args.channelIconURL, actual.ChannelIconURL)
			assert.EqualValues(t, tt.args.channelURL, actual.ChannelURL)
			assert.EqualValues(t, tt.args.fab, actual.Fab)
			assert.EqualValues(t, tt.args.createdAt, actual.CreatedAt)
			assert.EqualValues(t, tt.args.updatedAt, actual.UpdatedAt)
		})
	}
}

func TestVideo_Favorite(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		v    *Video
	}{
		{
			name: "success",
			v:    &Video{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.v.Favorite()
			assert.True(t, tt.v.Fab, "Favorite should set Fab to true")
		})
	}
}

func TestVideo_Unfavorite(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		v    *Video
	}{
		{
			name: "success",
			v:    &Video{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.v.Unfavorite()
			assert.False(t, tt.v.Fab, "Unfavorite should set Fab to false")
		})
	}
}

func TestVideo_String(t *testing.T) {
	t.Parallel()

	type fields struct {
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
		DeletedAt      *time.Time
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				ID:             "id",
				VideoID:        "videoID",
				Title:          "title",
				Description:    "description",
				ChannelName:    "channelName",
				ChannelIconURL: "channelIconURL",
				ChannelURL:     "channelURL",
				Fab:            true,
				CreatedAt:      time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:      time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				DeletedAt:      nil,
			},
			want: "ID : id\tVideoID : videoID\tTitle : title\tDescription : description\tChannelName : channelName\tChannelIconURL : channelIconURL\tChannelURL : channelURL\tFab : true\tCreatedAt : 2020-01-01 00:00:00\tUpdatedAt : 2020-01-01 00:00:00\tDeletedAt : <empty>",
		},
		{
			name: "success",
			fields: fields{
				ID:             "id",
				VideoID:        "videoID",
				Title:          "title",
				Description:    "description",
				ChannelName:    "channelName",
				ChannelIconURL: "channelIconURL",
				ChannelURL:     "channelURL",
				Fab:            true,
				CreatedAt:      time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:      time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				DeletedAt:      PointerTime(t),
			},
			want: "ID : id\tVideoID : videoID\tTitle : title\tDescription : description\tChannelName : channelName\tChannelIconURL : channelIconURL\tChannelURL : channelURL\tFab : true\tCreatedAt : 2020-01-01 00:00:00\tUpdatedAt : 2020-01-01 00:00:00\tDeletedAt : 2020-01-01 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			v := &Video{
				ID:             tt.fields.ID,
				VideoID:        tt.fields.VideoID,
				Title:          tt.fields.Title,
				Description:    tt.fields.Description,
				ChannelName:    tt.fields.ChannelName,
				ChannelIconURL: tt.fields.ChannelIconURL,
				ChannelURL:     tt.fields.ChannelURL,
				Fab:            tt.fields.Fab,
				CreatedAt:      tt.fields.CreatedAt,
				UpdatedAt:      tt.fields.UpdatedAt,
				DeletedAt:      tt.fields.DeletedAt,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("Video.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
