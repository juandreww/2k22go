package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://clara:password@localhost/employees?sslmode=disable")
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)

	fmt.Println("Welcome to the postgres.")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
