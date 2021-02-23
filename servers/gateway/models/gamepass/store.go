package gamepass

//Store interface for password, very insecure by design
type Store interface {
	Compare(password string) (string, error)
	UpdateTrackable(password string) error
	UpdateNonTrackable(password string) error
	Get() (string, string, error)
}
