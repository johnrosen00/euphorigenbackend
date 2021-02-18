package gamepass

//Store interface for password
type Store interface {
	UpdateTrackable(password string) error
	UpdateNonTrackable(password string) error
}
