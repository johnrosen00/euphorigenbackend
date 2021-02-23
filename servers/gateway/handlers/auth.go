package handlers

import (
	"encoding/json"
	"euphorigenbackend/servers/gateway/sessions"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//returns true if current request header = application/json
func isJSONctype(r *http.Request) bool {
	return r.Header.Get("Content-Type") == "application/json"
}

//Credentials = login credentials
type Credentials struct {
	Password string `json:"password"`
}

//SessionsHandler creates new sesssions
func (cx *HandlerContext) SessionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Forbidden method.", http.StatusMethodNotAllowed)
		return
	}

	if !isJSONctype(r) {
		http.Error(w, "Request body must contain JSON.", http.StatusUnsupportedMediaType)
		return
	}

	//	request body must contain json that can be decoded into a users.Credentials struct
	//	if:
	//	- cannot find user using credentials OR
	//	- invalid json format OR
	//	- can find user using credentials but cannot authenticate.
	//		http.Error(w, "invalid credentials", http.StatusUnauthorized)

	newCredentials := &Credentials{}
	d := json.NewDecoder(r.Body)

	if err := d.Decode(newCredentials); err != nil {
		bcrypt.CompareHashAndPassword([]byte("g"), []byte(newCredentials.Password))
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	if bcrypt.CompareHashAndPassword(cx.ManPass, []byte(newCredentials.Password)) != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	//	start new session.
	newSessionState := &SessionState{}
	newSessionState.StartTime = time.Now()
	sessions.BeginSession(cx.Key, cx.SessionStore, newSessionState, w)

	//	response body:
	//		content-type header = application/json
	// 		status code http.StatusCreated
	//		json encoded copy of user's profile
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
