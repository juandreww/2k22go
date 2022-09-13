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

    
    mux := http.DefaultServeMux
    mux.HandleFunc("/", index)
    mux.HandleFunc("/whatshouldiwear", whatshouldiwear)
    mux.HandleFunc("/create", create)
    mux.HandleFunc("/read", read)
    mux.HandleFunc("/update", update)
    mux.HandleFunc("/delete", delete)
    mux.HandleFunc("/drop", drop)
    mux.Handle("/favicon.ico", http.NotFoundHandler())
    err = http.ListenAndServe(":8080", nil)
    check(err)
}

func index(w http.ResponseWriter, r *http.Request) {
    _, err = io.WriteString(w, "Successfully Completed")
    check(err)
}

func whatshouldiwear(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, name FROM default.basicdata;")
    check(err)
    defer rows.Close()

    var s, name string
    s = "RETRIEVED RECORDS:\n"

    for rows.Next() {
        err = rows.Scan(&name)
        check(err)
        s += name + "\n"
    }
}

func create(w http.ResponseWriter, r *http.Request) {
    stmt, err := db.Prepare(`CREATE TABLE customer(id int, name varchar(20));`)
    check(err)
    defer stmt.Close()

    row, err := stmt.Exec()
    check(err)

    n, err := row.RowsAffected()
    check(err)

    fmt.Fprintln(w, "CREATED TABLE customer", n)
}

func insert(w http.ResponseWriter, r *http.Request) {
    stmt, err := db.Prepare(`INSERT INTO customer VALUES ("Hayabusa"), ("Doris Doe");`)
    check(err)
    defer stmt.Close()

    row, err := stmt.Exec()
    check(err)

    n, err := row.RowsAffected()
    check(err)

    fmt.Println(w, "INSERTED NEW RECORD", n)
}

func read(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query(`SELECT * FROM customer;`)
    check(err)
    defer rows.Close()

    var name string
    for rows.Next() {
        err = rows.Scan(&name)
        check(err)
        fmt.Fprintln(w, "RETRIEVED RECORD:", name)
    }
}

func update(w http.ResponseWriter, r *http.Request) {
    stmt, err := db.Prepare(`UPDATE customer SET name="Jack Whales" WHERE name="Hayabusa"`)
    check(err)
    defer stmt.Close()

    row, err := stmt.Exec()
    check(err)

    n, err := row.RowsAffected()
    check(err)

    fmt.Fprintln(w, "UPDATED RECORD:", n)
}

func delete(w http.ResponseWriter, r *http.Request) {
    stmt, err := db.Prepare(`DELETE FROM customer WHERE name="Jack Whales";`)
    check(err)
    defer stmt.Close()

    row, err := stmt.Exec()
    check(err)

    n, err := row.RowsAffected()
    check(err)

    fmt.Fprintln(w, "DELETED RECORD:", n)
}

func drop(w http.ResponseWriter, r *http.Request) {
    stmt, err := db.Prepare(`DROP TABLE customer;`)
    check(err)
    defer stmt.Close()

    _, err = stmt.Exec()
    check(err)

    fmt.Fprintln(w, "DROPPED TABLE customer")
}

func check(err error) {
    if err != nil {
        panic(err)
    } 
}

