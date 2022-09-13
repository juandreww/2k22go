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
    mux := http.DefaultServeMux
    mux.HandleFunc("/index", index)
    err := http.ListenAndServe(":8080", nil)

    check(err)
    
}

func check(err error) {
    if err != nil {
        panic(err)
    } 
}

