package gamepass

import "database/sql"

//PassStore stores the password stuff
type PassStore struct {
	DB *sql.DB
}
