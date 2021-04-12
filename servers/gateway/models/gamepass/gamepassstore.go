package gamepass

import (
	"database/sql"
	"fmt"
)

//PassStore stores the password stuff
type PassStore struct {
	DB *sql.DB
}

//checks to see if there exactly 1 entry in the password table
//If there is not exactly 1 row, sets passwords to "garfield", "garfield2"
func (store *PassStore) checkState() {
	q := "select tpass, ntpass from passwords"
	rows, err := store.DB.Query(q)

	if err != nil {
		return
	}

	i := 0
	for rows.Next() {
		i++
	}

	if i < 1 {
		//delete all rows, set to default, throw error
		q3 := "insert into passwords(tpass, ntpass) values(?,?)"
		_, err = store.DB.Exec(q3, "garfield", "garfield2")
		if err != nil {
			fmt.Printf("error inserting new row: %v\n", err)
			return
		}
	}
	return
}

//Compare compares param password to password stored in DB.
//returns "track", "notrack"
func (store *PassStore) Compare(password string) (string, error) {
	ret := ""
	store.checkState()
	q := "select tpass, ntpass from passwords"
	row := store.DB.QueryRow(q)

	track := ""
	notrack := ""

	if err := row.Scan(&track, &notrack); err != nil {
		return "", err
	}

	if password == track {
		ret = "track"
	} else if password == notrack {
		ret = "notrack"
	}

	return ret, nil
}

//UpdateTrackable updates the trackable password
func (store *PassStore) UpdateTrackable(password string) error {
	store.checkState()
	queryF := "update passwords set tpass = ?"

	if _, err := store.DB.Exec(queryF, password); err != nil {
		return err
	}

	return nil
}

//UpdateNonTrackable updates the nontrackable password.
func (store *PassStore) UpdateNonTrackable(password string) error {
	store.checkState()
	queryF := "update passwords set ntpass = ?"

	if _, err := store.DB.Exec(queryF, password); err != nil {
		return err
	}

	return nil
}

//Get returns the nontrackable and trackable passwords
func (store *PassStore) Get() (string, string, error) {
	store.checkState()

	q := "select tpass, ntpass from passwords"
	row := store.DB.QueryRow(q)

	track := ""
	notrack := ""

	if err := row.Scan(&track, &notrack); err != nil {
		return "", "", err
	}

	return track, notrack, nil
}
