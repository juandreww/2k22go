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
	rows, err := db.Query("SELECT * FROM pricing;")
	checkError(err)
	defer rows.Close()

	prices := make([]Pricing, 0)
	for rows.Next() {
		pc := Pricing{}
		err := rows.Scan(&pc.ID, &pc.Title, &pc.Price)
		checkError(err)
		prices = append(prices, pc)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
