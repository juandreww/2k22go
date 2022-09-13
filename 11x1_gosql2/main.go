package main

import (
    "database/sql"
    // "fmt"
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

    
    mux := http.DefaultServeMux
    mux.HandleFunc("/index", index)
    mux.Handle("/favicon.ico", http.NotFoundHandler())
    err = http.ListenAndServe(":8080", nil)
    check(err)
}

func index(w http.ResponseWriter, r *http.Request) {
    _, err = io.WriteString(w, "Successfully Completed")
    check(err)
}

func check(err error) {
    if err != nil {
        panic(err)
    } 
}

