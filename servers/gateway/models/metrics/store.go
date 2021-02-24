package metrics

//Store of Metrics
type Store interface {
	UpdateTrackable(password string) error
	UpdateNonTrackable(password string) error
	Get() ([]*Metric, error)
}
