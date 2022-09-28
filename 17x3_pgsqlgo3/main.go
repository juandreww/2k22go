package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Pricing struct {
	ID    string
	Title string
	Price float32
}

var db *sql.DB
var tpl *template.Template

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://clara:password@localhost/cube0?sslmode=disable")
	checkError(err)

	err = db.Ping()
	checkError(err)

	fmt.Println("Welcome to the postgres.")
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/index/show", indexShow)
	http.HandleFunc("/index/create", indexCreateForm)
	http.HandleFunc("/index/create/process", indexCreateProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

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

	tpl.ExecuteTemplate(w, "prices.gohtml", prices)
}

func indexShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM pricing WHERE id = $1;", id)
	pc := Pricing{}
	err := row.Scan(&pc.ID, &pc.Title, &pc.Price)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "show.gohtml", pc)
}

func indexCreateForm(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "create.gohtml", nil)
}

func indexCreateProcess(w http.ResponseWriter, r *http.Request) {

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
