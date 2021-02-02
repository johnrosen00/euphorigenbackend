package users

import (
	"database/sql"
	"fmt"

	//comment to satisfy lintr
	_ "github.com/go-sql-driver/mysql"
)

//MySQLStore stores a pointer to a database
type MySQLStore struct {
	DB *sql.DB
}

//GetByID returns the User with the given ID
func (store *MySQLStore) GetByID(id int64) (*User, error) {
	q := "select userid, email, passhash from users where id = ?"
	row := store.DB.QueryRow(q, id)

	u := &User{}
	err := row.Scan(&u.ID, &u.Email, &u.PassHash)

	if err != nil {
		return nil, err
	}

	return u, nil
}

//GetByEmail returns the User with the given email
func (store *MySQLStore) GetByEmail(email string) (*User, error) {
	q := "select userid, email, passhash from users where email = ?"
	row := store.DB.QueryRow(q, email)

	u := &User{}
	err := row.Scan(&u.ID, &u.Email, &u.PassHash)

	if err != nil {
		return nil, err
	}

	return u, nil
}

//Insert inserts the user into the database, and returns
//the newly-inserted User, complete with the DBMS-assigned ID
func (store *MySQLStore) Insert(user *User) (*User, error) {
	insq := "insert into users(email, passhash) values(?,?)"
	//INSert Query
	res, err := store.DB.Exec(insq, user.Email, user.PassHash)
	if err != nil {
		fmt.Printf("error inserting new row: %v\n", err)
		return nil, err
	}

	//get the auto-assigned ID for the new row
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("error getting new ID: %v\n", id)
		return nil, err
	}

	return user, nil
}

//Delete deletes the user with the given ID
func (store *MySQLStore) Delete(id int64) error {
	ex := "delete from users where userid = ?"

	_, err := store.DB.Exec(ex, id)

	if err != nil {
		return err
	}

	return nil
}
