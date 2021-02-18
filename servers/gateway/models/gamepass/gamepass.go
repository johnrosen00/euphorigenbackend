package gamepass

//NewPass is used when setting new password
type NewPass struct {
	Password string `json:"Password"`
}

//PlayerLogin is used when players try to login to an existing session.
//maybe not going to use this
type PlayerLogin struct {
	Password []byte `json:"Password"`
	ID       int64  `json:"ID"`
}
