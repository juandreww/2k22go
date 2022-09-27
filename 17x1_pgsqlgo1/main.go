package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Pricing struct {
	ID    string
	Title string
	Price string
}

func main() {
	db, err := sql.Open("postgres", "postgres://clara:password@localhost/cube0?sslmode=disable")
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
