package players

//Store interface for password
type Store interface {
	Insert() (int64, error)
	Update(id int64, newpuzzle int64) error
	Validate(id int64) error
}
