package handlers

import (
	"euphorigenbackend/servers/gateway/sessions"
)

//HandlerContext contains handler context
type HandlerContext struct {
	Key          string
	SessionStore sessions.Store
	ManPass      []byte
}

//GameContext is a struct that is used for handlers for the nonmanagement endpoints
type GameContext struct {
	Key          string
	SessionStore sessions.Store
	PassStore    string
}
