package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
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
	http.HandleFunc("/index/update", indexUpdateForm)
	http.HandleFunc("/index/update/process", indexUpdateProcess)
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
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	bk := Pricing{}
	bk.ID = r.FormValue("id")
	bk.Title = r.FormValue("title")
	p := r.FormValue("price")

	// validate form values
	if bk.ID == "" || bk.Title == "" || p == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please hit back and enter a number for the price", http.StatusNotAcceptable)
		return
	}
	bk.Price = float32(f64)

	// insert values
	_, err = db.Exec("INSERT INTO pricing (id, title, price) VALUES ($1, $2, $3)", bk.ID, bk.Title, bk.Price)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// confirm insertion
	tpl.ExecuteTemplate(w, "created.gohtml", bk)
}

func indexUpdateForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM pricing WHERE id = $1", id)

	bk := Pricing{}
	err := row.Scan(&bk.ID, &bk.Title, &bk.Price)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "update.gohtml", bk)
}

func indexUpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	bk := Pricing{}
	bk.ID = r.FormValue("id")
	bk.Title = r.FormValue("title")
	p := r.FormValue("price")

	// validate form values
	if bk.ID == "" || bk.Title == "" || p == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please hit back and enter a number for the price", http.StatusNotAcceptable)
		return
	}
	bk.Price = float32(f64)

	// insert values
	_, err = db.Exec("UPDATE pricing SET id = $1, title=$2, price=$3 WHERE id=$1;", bk.ID, bk.Title, bk.Price)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// confirm insertion
	tpl.ExecuteTemplate(w, "updated.gohtml", bk)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
