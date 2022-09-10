package main

import (
	"fmt"
	"html/template"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
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

	i := 0;
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

		list[i] = currency{
			value.ID,
			value.Name,
		}
		i++
	}
	fmt.Println(list)

	tpl.ExecuteTemplate(w, "listcurrency.gohtml", nil)
}