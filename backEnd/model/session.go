package model

import "time"

type Session struct {
	UserID    int    `json:"user_id"`
	SessionID string    `json:"session_id"`
	ExpiredAt time.Time `json:"expired_at"`
}
