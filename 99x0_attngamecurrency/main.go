package main

import (
	"fmt"
	"html/template"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	// "log"
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
	sqlStatement := `SELECT id, name FROM currency WHERE id=$1;`
	row := con.QueryRow(sqlStatement, data.ID)
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
		tpl.ExecuteTemplate(w, "index.gohtml", data)
	} else {
		data = currency{
			"error",
			"error",
		}
	}
	

	
}