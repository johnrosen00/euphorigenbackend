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
	q := "select * from metrics"
	rows, err := store.DB.Query(q)
	if err != nil {
		return nil, err
	}

	var metricSlice []*Metric
	var m *Metric
	defaultTime := time.Time{}
	for rows.Next() {

		if err = rows.Scan(&m.MetricID, &m.PlayerID, &m.PuzzleID, &m.TimeInitiated); err != nil {
			return nil, err
		}

		if mr.BeginTime != defaultTime && mr.EndTime != defaultTime {
			if m.TimeInitiated.Before(mr.EndTime) && m.TimeInitiated.After(mr.BeginTime) {
				//do nothing
			} else {
				continue
			}
		} else if mr.BeginTime != defaultTime {
			if !m.TimeInitiated.After(mr.BeginTime) {
				continue
			}

		} else if mr.EndTime != defaultTime {
			if !m.TimeInitiated.Before(mr.EndTime) {
				//do nothing
			} else {
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

	return metricSlice, nil
}
