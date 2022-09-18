package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "aws_procube1:aws_procube1@tcp(aws-procube1.chjuuutofqif.ap-southeast-3.rds.amazonaws.com:3306)/default?charset=utf8")
    check(err)
    defer db.Close()

    err = db.Ping()
    check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.HandleFunc("/retrievedata", retrieveData)
    http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from AWS.")
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
}

func retrieveData(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT id, name FROM default.basicdata;`)
	check(err)

	// data to be used in query
	s := getInstance()
	s += "\nRETRIEVED RECORDS:\n"
	var id, name string

	// query
	for rows.Next() {
		err = rows.Scan(&id, &name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func instance(w http.ResponseWriter, req *http.Request) {
	s := getInstance()
	io.WriteString(w, s)
}


func getInstance() string {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	resp.Body.Close()
	return string(bs)
}

func check(err error) {
    if err != nil {
        panic(err)
    } 
}
