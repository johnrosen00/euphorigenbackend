package metrics

//Store of Metrics
type Store interface {
	Insert(m *Metric) error
	Get(mr *MetricRequest) ([]*Metric, error)
}
