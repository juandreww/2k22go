package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type Pricing struct {
	ID    string
	Title string
	Price float32
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://clara:password@localhost/cube0?sslmode=disable")
	checkError(err)

	err = db.Ping()
	checkError(err)

	fmt.Println("Welcome to the postgres.")
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "GET" {
	// 	http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	// 	return
	// }

	rows, err := db.Query("SELECT * FROM pricing;")
	if err != nil {
		fmt.Println("44")
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	prices := make([]Pricing, 0)
	for rows.Next() {
		pc := Pricing{}
		err := rows.Scan(&pc.ID, &pc.Title, &pc.Price)
		if err != nil {
			fmt.Println("54")
			http.Error(w, http.StatusText(500), 500)
			return
		}
		prices = append(prices, pc)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("63")
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, pc := range prices {
		fmt.Printf("%s, %s, %.2f\n", pc.ID, pc.Title, pc.Price)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
