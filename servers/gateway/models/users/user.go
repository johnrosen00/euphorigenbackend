package users

import (
	"fmt"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

var bcryptCost = 13

//Credentials: login credentials
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//NewUser: struct created when new user created
type NewUser struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	PasswordConf   string `json:"passwordconf"`
	ServerPassword string `json:"serverpassword"`
}

//User: User struct for sessions.
type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"-"` //never JSON encoded/decoded
	PassHash []byte `json:"-"` //never JSON encoded/decoded
}

//Validate: validates new user struct
func (nu *NewUser) Validate() error {

	if _, err := mail.ParseAddress(nu.Email); err != nil {
		return fmt.Errorf("BadEmailFormat")
	}

	//check env variable

	return nil
}

//ToUser converts the NewUser to a User
func (nu *NewUser) ToUser() (*User, error) {
	if err := nu.Validate(); err != nil {
		return nil, err
	}
	u := &User{}

	u.ID = 0

	if err1 := u.SetPassword(nu.Password); err1 != nil {
		return nil, err1
	}

	return u, nil
}

//SetPassword hashes the password and stores it in the PassHash field
func (u *User) SetPassword(password string) error {
	//use the bcrypt package to generate a new hash of the password
	//https://godoc.org/golang.org/x/crypto/bcrypt

	bhash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)

	if err != nil {
		return err
	}

	u.PassHash = bhash

	return nil
}

//Authenticate compares the plaintext password against the stored hash
//and returns an error if they don't match, or nil if they do
func (u *User) Authenticate(password string) error {
	//use the bcrypt package to compare the supplied
	//password with the stored PassHash
	//https://godoc.org/golang.org/x/crypto/bcrypt

	return bcrypt.CompareHashAndPassword(u.PassHash, []byte(password))
}
