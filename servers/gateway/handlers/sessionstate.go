package handlers

import (
	"time"
)

//SessionState describes current user session
type SessionState struct {
	StartTime       time.Time `json:"starttime"`
	PlayerSessionID int64     `json:"playerSessionID"`
}
