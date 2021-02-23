package handlers

import (
	"time"
)

//SessionState describes current user session
type SessionState struct {
	StartTime time.Time `json:"starttime"`
}

//PlayerSessionState contains a player session
type PlayerSessionState struct {
	StartTime       time.Time `json:"startTime"`
	PlayerSessionID int64     `json:"playerSessionID"`
}
