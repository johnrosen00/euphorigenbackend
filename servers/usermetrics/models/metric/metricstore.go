package metrics

import (
	"database/sql"
	"fmt"
	"time"
)

//MetricStore implements store interface, acts as a store of metrics
type MetricStore struct {
	DB *sql.DB
}

//GetMetricTypeByID returns metrictype by ID. Wild.
func (store *MetricStore) GetMetricTypeByID(id int64) (*MetricType, error) {
	q := "select metrictypeid, metrictypename from users where metrictypeid = ?"
	row := store.DB.QueryRow(q, id)

	mt := &MetricType{}
	err := row.Scan(&mt.MetricTypeID, &mt.MetricTypeName)

	if err != nil {
		return nil, err
	}

	return mt, nil
}

//GetMetricTypeByName returns metrictype by name. Wild.
func (store *MetricStore) GetMetricTypeByName(name string) (*MetricType, error) {
	q := "select metrictypeid, metrictypename from users where metrictypename = ?"
	row := store.DB.QueryRow(q, name)

	mt := &MetricType{}
	err := row.Scan(&mt.MetricTypeID, &mt.MetricTypeName)

	if err != nil {
		return nil, err
	}

	return mt, nil
}

//GetMetrics description:
// Params: params []byte
// params is a json file that contains the request body of /v1/game/manage/metrics
// might just change this to a struct, I don't know.
// returns a list of metrics that follow specified params
func (store *MetricStore) GetMetrics(params []byte) ([]*Metric, error) {
	return nil, nil
}

//InsertMetric inserts a metric into the database
func (store *MetricStore) InsertMetric(metric *Metric) (*Metric, error) {
	err := store.Validate(metric)

	if err != nil {
		return nil, err
	}

	newMetric := metric
	return newMetric, nil
}

//Validate checks to see whether metric posted is valid
func (store *MetricStore) Validate(m *Metric) error {
	if m.Timestamp.After(time.Now()) {
		return fmt.Errorf("Invalid timestamp")
	}

	if _, err := store.GetMetricTypeByID(m.MetricTypeID); err != nil {
		return fmt.Errorf("Invalid metrictype")
	}

	return nil
}
