package metrics

import "database/sql"

//MetricStore is just a way to add stuff to the db
type MetricStore struct {
	DB *sql.DB
}

//Insert a new metric into the DB
func (store *MetricStore) Insert(m *Metric) error {
	return nil
}

//Get events based on MetricRequest params
func (store *MetricStore) Get(mr *MetricRequest) ([]*Metric, error) {
	return nil, nil
}
