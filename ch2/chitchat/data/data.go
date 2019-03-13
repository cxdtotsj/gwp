package data

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=bear sslmode=disabel")
	if err != nil {
		log.Fatal(err)
	}
	return
}
