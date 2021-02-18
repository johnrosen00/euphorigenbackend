package handlers

import "net/http"

//PlayerHandler handles player login
func (gcx *GameContext) PlayerHandler(w http.ResponseWriter, r *http.Request) {

	// v1/game/session
	if r.Method != "POST" {
		http.Error(w, "Forbidden method.", http.StatusMethodNotAllowed)
		return
	}

}
