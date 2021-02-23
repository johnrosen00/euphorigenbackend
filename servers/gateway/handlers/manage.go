package handlers

import (
	"encoding/json"
	"euphorigenbackend/servers/gateway/models/gamepass"
	"euphorigenbackend/servers/gateway/sessions"
	"net/http"
)

//manages gamepass

//GamePassHandler allows manager to modify game password
//v1/game/manage/password
func (cx *HandlerContext) GamePassHandler(w http.ResponseWriter, r *http.Request) {

	//Authorize
	currentSessionState := &SessionState{}
	if _, errGetSession := sessions.GetState(r, cx.Key, cx.SessionStore, currentSessionState); errGetSession != nil {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}
	if currentSessionState.PlayerSessionID != -1 {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}

	if r.Method == "POST" {
		if !IsJSONctype(r) {
			http.Error(w, "Request body must contain JSON.", http.StatusUnsupportedMediaType)
			return
		}
		//get the request body
		newpass := &gamepass.NewPass{}
		d := json.NewDecoder(r.Body)
		if err := d.Decode(newpass); err != nil {
			http.Error(w, "Bad JSON body", http.StatusUnsupportedMediaType)
			return
		}

		if newpass.Track {
			cx.GamePassStore.UpdateTrackable(newpass.Password)
		} else {
			cx.GamePassStore.UpdateNonTrackable(newpass.Password)
		}
	} else if r.Method == "GET" {
		t, nt, err := cx.GamePassStore.Get()
		if err != nil {
			http.Error(w, "Bad response from server", http.StatusInternalServerError)
			return
		}
		ret := &struct {
			TPassword  string `json:"TPassword"`
			NTPassword string `json:"NTPassword"`
		}{
			t, nt,
		} //cheeky anonymized struct

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)

		if err3 := enc.Encode(ret); err3 != nil {
			http.Error(w, "Unable to encode to JSON", 404)
		}

	} else {
		http.Error(w, "Request body must contain JSON.", http.StatusMethodNotAllowed)
		return

	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}