package handlers

import (
	"euphorigenbackend/servers/gateway/models/gamepass"
	"euphorigenbackend/servers/gateway/models/metrics"
	"euphorigenbackend/servers/gateway/models/players"

	"euphorigenbackend/servers/gateway/sessions"
)

//HandlerContext contains handler context
type HandlerContext struct {
	Key           string
	SessionStore  sessions.Store
	ManPass       []byte
	GamePassStore gamepass.Store
	PlayerStore   players.Store
	MetricStore   metrics.Store
}
