package handlers

import (
	"euphorigenbackend/servers/gateway/models/users"
	"euphorigenbackend/servers/gateway/sessions"
)

//HandlerContext contains handler context
type HandlerContext struct {
	Key          string
	SessionStore sessions.Store
	UserStore    users.Store
}

//GameContext is a struct that is used for handlers for the nonmanagement endpoints
type GameContext struct {
	Key          string
	SessionStore sessions.Store
	PassStore    string
}
