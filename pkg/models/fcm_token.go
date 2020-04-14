package models

import (
	"errors"
	"time"
)

// Typed errors
var (
	FCMFieldMissing    = errors.New("missing token or user Id")
	ErrTokenRegistered = errors.New("this token is already regstered...")
)

// FCM token model
type FcmToken struct {
	Id     int64
	Token  string
	UserId string

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// ---------------------
// COMMANDS
type AddTokenCommand struct {
	Token  string
	UserId string

	Result FcmToken `json:"-"`
}

// ---------------------
// QUERIES

type GetTokenByUserQuery struct {
	UserId []string
	Result []string
}

type GetTokeQuery struct {
	Result []*FcmToken `json:"-"`
}
