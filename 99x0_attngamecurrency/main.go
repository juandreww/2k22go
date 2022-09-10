package main

import (
	"fmt"
	"html/template"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	// "os"
)

const (
	host = "ec2-52-73-155-171.compute-1.amazonaws.com"
	port = 5432
	user = "'awtpbzyctlpydm'"
	password = "a7bf40c39496f73a03e7412befbc787d29138445d7fce2a34bf31df40cf07d96"
	dbname = "d4ehughfapgq0k"
)

type currency struct {
	ID string
	Name string
}

type configconvertrate struct {
	CurrencyFrom string
	CurrencyTo string
	Rate string
}

var tpl *template.Template
var con *sql.DB

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=require",
	host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	con = db
	return db
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	mux.HandleFunc("/savecurrency", savecurrency)
	mux.HandleFunc("/listcurrency", listcurrency)
	mux.HandleFunc("/addconversionrate", addconversionrate)
	db := ConnectDB()
	defer db.Close()

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is index api: ", r.Method)

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func savecurrency(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is savecurrency api: ", r.Method)

	data := currency{
		r.FormValue("id"),
		r.FormValue("name"),
	}

	fmt.Println(data)

	value := currency{}
	isexist := false
	sqlStatement := `SELECT id, name FROM currency WHERE (id=$1 OR name=$2);`
	row := con.QueryRow(sqlStatement, data.ID, data.Name)
	err := row.Scan(&value.ID, &value.Name,)
	switch err {
	case sql.ErrNoRows:
		isexist = false
	case nil:
		isexist = true
	default:
	  	panic(err)
	}
	
	if isexist == false {
		sqlStatement = `
			INSERT INTO currency (id, name)
			VALUES ($1, $2)`
		_, err := con.Exec(sqlStatement, data.ID, data.Name)
		if err != nil {
			panic(err)
		}
	} else {
		data = currency{
			"error",
			"error",
		}
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func listcurrency(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is listcurrency api: ", r.Method)
	var list []currency

	rows, err := con.Query("SELECT id, name FROM currency")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		value := currency{}
		err := rows.Scan(&value.ID, &value.Name,)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("row is not exist")
			return
		case nil:
			fmt.Println(value.ID, value.Name)
		default:
			panic(err)
		}

		list = append(list, value)
	}

	tpl.ExecuteTemplate(w, "listcurrency.gohtml", list)
}

func addconversionrate(w http.ResponseWriter, r *http.Request) {
	if (r.Method == http.MethodPost) {
		fmt.Println("This is add conversion rate api: ", r.Method)
		data := configconvertrate{
			r.FormValue("currencyfrom"),
			r.FormValue("currencyto"),
			r.FormValue("rate"),
		}
	
		fmt.Println(data)

		var value string
		var intval int
		check1 := currency{}

		sqlStatement := `SELECT count(id) id FROM currency WHERE (id=$1 OR id=$2);`
		row := con.QueryRow(sqlStatement, data.CurrencyFrom)
		err := row.Scan(&check1.ID, &check1.Name)
		isError := HandleErrorOfSelect(w, err)
		if (isError == true) {
			tmp := currency{
				"error",
				"Currency is not found in database",
			}
			tpl.ExecuteTemplate(w, "addconversionrate.gohtml", tmp)
			return
		}

		sqlStatement = `SELECT nullif(max(id),0) id FROM currencyrate`
		row = con.QueryRow(sqlStatement)
		if row == nil {
			fmt.Println("Row is nil")
			return
		}
		err = row.Scan(&value)
		isError = HandleErrorOfSelect(w, err)
		if (isError == true) {
			tmp := currency{
				"error",
				"CurrencyRate is not found in database",
			}
			tpl.ExecuteTemplate(w, "addconversionrate.gohtml", tmp)
			return
		}

		intval, err = strconv.Atoi(value)
		intval++
		sqlStatement = `
			INSERT INTO currencyrate (id, currencyfrom, currencyto, rate)
			VALUES ($1, $2, $3, $4)`
		_, err = con.Exec(sqlStatement, intval, data.CurrencyFrom, data.CurrencyTo, data.Rate)
		if err != nil {
			panic(err)
		}

		tpl.ExecuteTemplate(w, "addconversionrate.gohtml", data)
	} else {
		fmt.Println("This is add conversion rate api: ", r.Method)

		tpl.ExecuteTemplate(w, "addconversionrate.gohtml", nil)
	}
	
}

func HandleErrorOfSelect(w http.ResponseWriter, err error) bool {
	data := false
	switch err {
	case sql.ErrNoRows:
		data = true
	case nil:
		data = false
	default:
		data = true
	}

	return data
}