package handlers

import (
	"encoding/json"
	"euphorigenbackend/servers/gateway/models/metrics"
	"euphorigenbackend/servers/gateway/models/players"
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
		m := &metrics.Metric{}
		d := json.NewDecoder(r.Body)
		if err := d.Decode(m); err != nil {
			http.Error(w, "Bad JSON body", http.StatusUnsupportedMediaType)
			return
		}

		ret, err := cx.MetricStore.Insert(m)

		if err != nil {
			http.Error(w, "Error with metric params", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)

		if err := enc.Encode(ret); err != nil {
			http.Error(w, "Unable to encode to JSON", 404)
		}
	} else if r.Method == "GET" && currentplayerid == -1 {
		//Get player metrics based on some params

		mr := &metrics.MetricRequest{}
		d := json.NewDecoder(r.Body)
		if err := d.Decode(mr); err != nil {
			http.Error(w, "Bad JSON body", http.StatusUnsupportedMediaType)
			return
		}

		ret, err := cx.MetricStore.Get(mr)

		if err != nil {
			http.Error(w, "Error with metric params", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		enc := json.NewEncoder(w)

		if err := enc.Encode(ret); err != nil {
			http.Error(w, "Unable to encode to JSON", 404)
		}
	} else if r.Method == "PATCH" && currentplayerid > 0 {
		//get the request body
		playerupdate := &players.Player{}
		d := json.NewDecoder(r.Body)
		if err := d.Decode(playerupdate); err != nil {
			http.Error(w, "Bad JSON body", http.StatusUnsupportedMediaType)
			return
		}
		//get response body
		if cx.PlayerStore.Update(currentplayerid, playerupdate.LastPuzzleID) != nil {
			http.Error(w, "Something went wrong...", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
}
