package metrics

//Store represents a store for Users
type Store interface {
	GetMetricTypeByID(id int64) (*MetricType, error)
	GetMetrics(params []byte) ([]*Metric, error)
	InsertMetric(metric *Metric) (*Metric, error)
}
