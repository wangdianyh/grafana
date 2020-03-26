package models

import (
	"errors"
	"time"
)

// Typed errors
var (
	ErrFCMNotFound  = errors.New("FCM not found")
	FCMFieldMissing = errors.New("missing token or channel Id")
)

// FCM token model
type FcmToken struct {
	Id        int64
	Token     string
	ChannelId string
	UserId    string

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// ---------------------
// COMMANDS
type AddTokenCommand struct {
	Token     string
	ChannelId string
	UserId    string

	Result FcmToken `json:"-"`
}

// ---------------------
// QUERIES
type GetTokenByChannelQuery struct {
	ChannelId string
	Result    []*FcmToken
}

type GetTokenByUserQuery struct {
	UserId []string
	Result []string
}
