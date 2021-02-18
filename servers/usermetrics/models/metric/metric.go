package metrics

import (
	"time"
)

//Metric contains information about something that is being tracked by the puzzle
type Metric struct {
	PlayerSessionID int64
	MetricTypeID    int64
	Timestamp       time.Time
	PuzzleID        int64
	Info            string
}

//MetricType explains how a metric should be processed.
type MetricType struct {
	MetricTypeID   int64
	MetricTypeName int64
}
