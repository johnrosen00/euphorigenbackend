package metrics

import (
	"database/sql"
	"fmt"
	"time"
)

//MetricStore is a wrapper struct to interact with database
type MetricStore struct {
	DB *sql.DB
}

//Insert a new metric into the DB, returns newly minted metric
func (store *MetricStore) Insert(m *Metric) (*Metric, error) {
	q := "insert into metrics(playerid, puzzleid, timeinitiated, metrictype, info) values(?,?,?,?,?)"

	res, err := store.DB.Exec(q, m.PlayerID, m.PuzzleID, m.TimeInitiated, m.MetricType, m.Info)

	if err != nil {
		fmt.Printf("error inserting new row: %v\n", err)
		return nil, err
	}

	//get the auto-assigned ID for the new row
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("error getting new ID: %v\n", id)
		return nil, err
	}

	m.MetricID = id
	return m, nil
}

//Get events based on MetricRequest params
func (store *MetricStore) Get(mr *MetricRequest) ([]*Metric, error) {
	fmt.Printf("made it here pt 1")
	q := "select metricid, playerid, puzzleid, timeinitiated, metrictype, info from metrics"
	rows, err := store.DB.Query(q)
	if err != nil {
		return nil, err
	}

	fmt.Printf("made it here pt 2")

	var metricSlice []*Metric
	var m *Metric
	defaultTime := time.Time{}

	for rows.Next() {
		err := rows.Scan(&m.MetricID, &m.PlayerID, &m.PuzzleID, &m.TimeInitiated, &m.MetricType, &m.Info)
		if err != nil {
			return nil, err
		}
		timeInitiated, _ := time.Parse(time.RFC3339, m.TimeInitiated)

		if mr.BeginTime != defaultTime && mr.EndTime != defaultTime {
			if timeInitiated.Before(mr.EndTime) && timeInitiated.After(mr.BeginTime) {
				if mr.EndTime.Before(mr.BeginTime) {
					continue
				}
			} else {
				continue
			}
		} else if mr.BeginTime != defaultTime {
			if !timeInitiated.After(mr.BeginTime) {
				continue
			}

		} else if mr.EndTime != defaultTime {
			if !timeInitiated.Before(mr.EndTime) {
				continue
			}
		}

		if mr.MetricType != "" {
			if m.MetricType != mr.MetricType {
				continue
			}
		}

		if mr.PuzzleID > 0 {
			if mr.PuzzleID != m.PuzzleID {
				continue
			}
		}

		metricSlice = append(metricSlice, m)
	}

	fmt.Printf("made it here pt 3")

	return metricSlice, nil
}
