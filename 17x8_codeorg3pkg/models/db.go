package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://clara:password@localhost/cube0?sslmode=disable")
	checkError(err)

	err = db.Ping()
	checkError(err)

	fmt.Println("Welcome to the postgres.")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
