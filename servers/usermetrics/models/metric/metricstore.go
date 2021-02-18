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
func (ms *MetricStore) GetMetricTypeByID(id int64) (*MetricType, error) {
	return nil, nil
}

//GetMetrics description:
// Params: params []byte
// params is a json file that contains the request body of /v1/game/manage/metrics
// might just change this to a struct, I don't know.
// returns a list of metrics that follow specified params
func (ms *MetricStore) GetMetrics(params []byte) ([]*Metric, error) {
	return nil, nil
}

//InsertMetric inserts a metric into the database
func (ms *MetricStore) InsertMetric(metric *Metric) (*Metric, error) {
	err := ms.Validate(metric)

	if err != nil {
		return nil, err
	}

	newMetric := metric
	return newMetric, nil
}

//Validate checks to see whether metric posted is valid
func (ms *MetricStore) Validate(m *Metric) error {
	if m.Timestamp.After(time.Now()) {
		return fmt.Errorf("Invalid timestamp")
	}

	if _, err := ms.GetMetricTypeByID(m.MetricTypeID); err != nil {
		return fmt.Errorf("Invalid metrictype")
	}

	return nil
}
