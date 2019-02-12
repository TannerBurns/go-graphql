package models

import (
	"database/sql"
	"fmt"
)

type Database struct {
	*sql.DB
}

// New makes a new database using the connection string and
// returns it, otherwise returns the error
func Session(host string, port int, database string, user string, password string) (session *Database, err error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database))
	if err != nil {
		return
	}
	// Check that our connection is good
	err = db.Ping()
	session = &Database{db}
	return
}
