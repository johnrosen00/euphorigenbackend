package metrics

import "time"

//Metric is a single metric
type Metric struct {
	PlayerID   int64     `json:"playerid"`
	PuzzleID   int64     `json:"puzzleid"`
	Timestamp  time.Time `json:"timestamp"`
	MetricType string    `json:"metrictype"`
	Info       string    `json:"info"`
}

//MetricRequest contains the parameters to GET many metrics
type MetricRequest struct {
	PlayerID   int64
	MetricType string
	BeginTime  time.Time
	EndTime    time.Time
	Info       string
} //don't know what other info should be capable of being retrieved?
