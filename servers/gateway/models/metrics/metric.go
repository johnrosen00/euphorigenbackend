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
	PlayerID   int64     `json:"playerid"`
	MetricType string    `json:"metrictype"`
	BeginTime  time.Time `json:"begintime"`
	EndTime    time.Time `json:"endtime"`
	Info       string    `json:"info"`
} //don't know what other info should be capable of being retrieved?
