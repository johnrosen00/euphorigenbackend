package handlers

import (
	"encoding/json"
	"euphorigenbackend/servers/gateway/models/metrics"
	"euphorigenbackend/servers/gateway/models/players"
	"euphorigenbackend/servers/gateway/sessions"
	"net/http"
	"strconv"
	"time"
)

type NewMetric struct {
	PuzzleID   int64  `json:"puzzleid"`
	MetricType string `json:"metrictype"`
	Info       string `json:"info"`
}

//MetricHandler handles the storage and retrieval of user metrics:
//POST
func (cx *HandlerContext) MetricHandler(w http.ResponseWriter, r *http.Request) {

	currentSessionState := &SessionState{}
	if _, errGetSession := sessions.GetState(r, cx.Key, cx.SessionStore, currentSessionState); errGetSession != nil {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}
	currentplayerid := currentSessionState.PlayerSessionID

	if r.Method == "POST" && currentplayerid > 0 {
		if !IsJSONctype(r) {
			http.Error(w, "Request body must contain JSON.", http.StatusUnsupportedMediaType)
			return
		}
		//Post a new metric
		nm := &NewMetric{}

		d := json.NewDecoder(r.Body)
		if err := d.Decode(nm); err != nil {
			http.Error(w, "Bad JSON body", http.StatusUnsupportedMediaType)
			return
		}

		m := &metrics.Metric{}

		m.Info = nm.Info
		m.PlayerID = currentplayerid
		m.MetricType = nm.MetricType
		m.MetricID = 0
		m.PuzzleID = nm.PuzzleID
		m.TimeInitiated = time.Now()
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
		mr.MetricType = r.FormValue("metrictype")

		puzzleParam := r.FormValue("puzzleid")

		if puzzleParam == "" {
			mr.PuzzleID = 0
		} else {
			puzzle, err := strconv.Atoi(puzzleParam)
			if err != nil {
				mr.PuzzleID = 0
			} else {
				mr.PuzzleID = int64(puzzle)
			}
		}

		beginTime := r.FormValue("begintime")
		t, err := time.Parse(time.RFC3339, beginTime)
		defaultTime := time.Time{}
		if err != nil {
			mr.BeginTime = defaultTime
		} else {
			mr.BeginTime = t
		}

		endTime := r.FormValue("endtime")
		t, err = time.Parse(time.RFC3339, endTime)

		if err != nil {
			mr.EndTime = defaultTime
		} else {
			mr.EndTime = t
		}

		ret, err := cx.MetricStore.Get(mr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		enc := json.NewEncoder(w)

		if err := enc.Encode(ret); err != nil {
			http.Error(w, "Unable to encode to JSON", 404)
		}
	} else if r.Method == "PATCH" && currentplayerid > 0 {
		if !IsJSONctype(r) {
			http.Error(w, "Request body must contain JSON.", http.StatusUnsupportedMediaType)
			return
		}
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
