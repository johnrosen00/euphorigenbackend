package players

import (
	"database/sql"
	"fmt"
)

//PlayerStore is a store of player session data
type PlayerStore struct {
	DB *sql.DB
}

//Insert creates a new player session
func (store *PlayerStore) Insert() (int64, error) {
	q := "insert into playersession(lastpuzzleid) values(?)"
	res, err := store.DB.Exec(q, 1)
	if err != nil {
		fmt.Printf("error inserting new row: %v\n", err)
		return -69, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("error getting new ID: %v\n", id)
		return -69, err
	}

	return id, nil
}

//Update updates a player session to reflect which puzzle they are currently on
func (store *PlayerStore) Update(id int64, newpuzzle int64) error {
	update := "update playersession set lastpuzzleid = ? where playerid = ?"

	if _, err := store.DB.Exec(update, newpuzzle, id); err != nil {
		return err
	}

	return nil
}
