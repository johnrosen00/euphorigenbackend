package handlers

import (
	"encoding/json"
	"euphorigenbackend/servers/gateway/sessions"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//IsJSONctype returns true if current request header = application/json
func IsJSONctype(r *http.Request) bool {
	return r.Header.Get("Content-Type") == "application/json"
}

//Credentials = login credentials
// Management = type of auth
type Credentials struct {
	Password   string `json:"password"`
	Management bool   `json:"management"`
}

//ReturnID struct contains a playerid
type ReturnID struct {
	ID int64 `json:"id"`
}

//SessionsHandler creates and gets new sesssions. PlayerID of -1 = management login, -2 = no-track
func (cx *HandlerContext) SessionsHandler(w http.ResponseWriter, r *http.Request) {

	if !IsJSONctype(r) {
		http.Error(w, "Request body must contain JSON.", http.StatusUnsupportedMediaType)
		return
	}
	//	request body must contain json that can be decoded into a users.Credentials struct
	//	if:
	//	- cannot find user using credentials OR
	//	- invalid json format OR
	//	- can find user using credentials but cannot authenticate.
	//		http.Error(w, "invalid credentials", http.StatusUnauthorized)
	ret := &ReturnID{}
	if r.Method == "POST" {
		newCredentials := &Credentials{}

		d := json.NewDecoder(r.Body)

		if err := d.Decode(newCredentials); err != nil {
			bcrypt.CompareHashAndPassword([]byte("g"), []byte(newCredentials.Password))
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		if newCredentials.Management {
			if bcrypt.CompareHashAndPassword(cx.ManPass, []byte(newCredentials.Password)) != nil {
				http.Error(w, "invalid credentials", http.StatusUnauthorized)
				return
			}

			//	start new session.
			newSessionState := &SessionState{}
			newSessionState.StartTime = time.Now()
			newSessionState.PlayerSessionID = -1
			ret.ID = -1
			sessions.BeginSession(cx.Key, cx.SessionStore, newSessionState, w)
		} else {
			pwtype, err := cx.GamePassStore.Compare(newCredentials.Password)
			if err != nil {
				http.Error(w, "invalid credentials", http.StatusUnauthorized)
				return
			}
			if pwtype == "track" {
				ret.ID, err = cx.PlayerStore.Insert()
				if err != nil {
					http.Error(w, "Database Error", http.StatusInternalServerError)
					return
				}
			} else if pwtype == "notrack" {
				ret.ID = -2
			} else {
				w.Header().Add("Content-Type", "application/json")
				http.Error(w, "Invalid Credentials.", http.StatusUnauthorized)
				return
			}

			newSessionState := &SessionState{}
			newSessionState.StartTime = time.Now()
			newSessionState.PlayerSessionID = ret.ID
			sessions.BeginSession(cx.Key, cx.SessionStore, newSessionState, w)
		}
	} else if r.Method == "GET" {

		currentSessionState := &SessionState{}
		if _, errGetSession := sessions.GetState(r, cx.Key, cx.SessionStore, currentSessionState); errGetSession != nil {
			http.Error(w, "Unauthorized user", http.StatusUnauthorized)
			return
		}

		ret.ID = currentSessionState.PlayerSessionID

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

	} else {
		http.Error(w, "Unsupported Method.", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)

	if err3 := enc.Encode(ret); err3 != nil {
		http.Error(w, "Unable to encode to JSON", 404)
	}
}
