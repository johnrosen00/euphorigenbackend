package metrics

import (
	"database/sql"
	"fmt"
)

//MetricStore is just a way to add stuff to the db
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

	var channelSlice []*Metric
	var m *Metric

	for rows.Next() {

		if err = rows.Scan(&m.MetricID, &m.PlayerID, &m.PuzzleID, &m.TimeInitiated); err != nil {
			return nil, err89
		}

		currentChannel, err45 := store.GetByID(currentChannelID)

		if err45 != nil {
			return nil, err45
		}

		channelSlice = append(channelSlice, currentChannel)
	}

	return channelSlice, nil
}
