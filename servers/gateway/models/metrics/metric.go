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

//MetricRequest contains the parameters to GET many metrics
type MetricRequest struct {
	PlayerID   int64     `json:"playerid,omitempty"`
	MetricType string    `json:"metrictype,omitempty"`
	BeginTime  time.Time `json:"begintime,omitempty"`
	EndTime    time.Time `json:"endtime,omitempty"`
	Info       string    `json:"info,omitempty"`
} //don't know what other info should be capable of being retrieved?
