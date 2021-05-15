package metrics

import "time"

//Metric is a single metric
type Metric struct {
	MetricID      int64     `json:"metricid"`
	PlayerID      int64     `json:"playerid"`
	PuzzleID      int64     `json:"puzzleid"`
	TimeInitiated time.Time `json:"timeinitiated"`
	MetricType    string    `json:"metrictype"`
	Info          string    `json:"info"`
}

type MetricRequest struct {
	MetricType string    `json:"metrictype,omitempty"`
	PuzzleID   int64     `json:"puzzleid"`
	BeginTime  time.Time `json:"begintime,omitempty"`
	EndTime    time.Time `json:"endtime,omitempty"`
}
