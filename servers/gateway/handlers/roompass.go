package handlers

import (
	"encoding/json"
	"euphorigenbackend/servers/gateway/sessions"
	"net/http"
)

type PuzzleInput struct {
	Password string `json:"password"`
	Puzzle   int    `json:"puzzle"`
}

func (cx *HandlerContext) RoomPassHandler(w http.ResponseWriter, r *http.Request) {
	//database = puzzle1
	puzzlePasswords := [5]string{"database", "inform", "fact", "9817", "media"}

	currentSessionState := &SessionState{}
	if _, errGetSession := sessions.GetState(r, cx.Key, cx.SessionStore, currentSessionState); errGetSession != nil {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}
	//compare passwords
	if r.Method == "POST" {
		//anonymized struct

		//get json input
		input := &PuzzleInput{}

		d := json.NewDecoder(r.Body)
		if err := d.Decode(input); err != nil {
			http.Error(w, "Bad JSON body", http.StatusUnsupportedMediaType)
			return
		}

		index := input.Puzzle - 1
		password := input.Password
		if index < 0 || index > len(puzzlePasswords)-1 {
			http.Error(w, "Invalid puzzle", http.StatusBadRequest)
		}

		w.Header().Add("Content-Type", "application/json")

		if puzzlePasswords[index] == password {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}

	} else {
		http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
	}

	return
}
