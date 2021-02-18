package handlers

import (
	"euphorigenbackend/servers/gateway/models/users"
	"time"
)

//SessionState describes current user session
type SessionState struct {
	StartTime time.Time   `json:"starttime"`
	User      *users.User `json:"user"`
}

//PlayerSessionState contains a player session
type PlayerSessionState struct {
	StartTime       time.Time `json:"startTime"`
	PlayerSessionID int64     `json:"playerSessionID"`
}
