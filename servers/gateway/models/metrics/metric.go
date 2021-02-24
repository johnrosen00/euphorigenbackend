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

//IncomingMetric is a thing because im not sure whether time.Time translates well into JSON
type IncomingMetric struct {
	MetricID      int64  `json:"metricid"`
	PlayerID      int64  `json:"playerid"`
	PuzzleID      int64  `json:"puzzleid"`
	TimeInitiated string `json:"timeinitiated"`
	MetricType    string `json:"metrictype"`
	Info          string `json:"info"`
}

//IncomingMetricRequest is a thing because i dont know whether time.Time translates well into JSON
type IncomingMetricRequest struct {
	PlayerID   int64
	MetricType string
	BeginTime  string
	EndTime    string
	Info       string
}

//MetricRequest contains the parameters to GET many metrics
type MetricRequest struct {
	PlayerID   int64
	MetricType string
	BeginTime  time.Time
	EndTime    time.Time
	Info       string
} //don't know what other info should be capable of being retrieved?
