package handlers

import (
	"euphorigenbackend/servers/gateway/sessions"
	"net/http"
)

//MetricHandler handles the storage and retrieval of user metrics:
//POST
func (cx *HandlerContext) MetricHandler(w http.ResponseWriter, r *http.Request) {
	if !IsJSONctype(r) {
		http.Error(w, "Request body must contain JSON.", http.StatusUnsupportedMediaType)
		return
	}

	currentSessionState := &SessionState{}
	if _, errGetSession := sessions.GetState(r, cx.Key, cx.SessionStore, currentSessionState); errGetSession != nil {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}
	currentplayerid := currentSessionState.PlayerSessionID

	if r.Method == "POST" && currentplayerid > 0 {
		//Post a new metric
	} else if r.Method == "GET" && currentplayerid == -1 {
		//Get player metrics based on some params
	} else if r.Method == "PATCH" && currentplayerid > 0 {
		//update current player puzzle
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
}
