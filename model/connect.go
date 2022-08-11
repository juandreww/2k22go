package model

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "awtpbzyctlpydm"
	password = "a7bf40c39496f73a03e7412befbc787d29138445d7fce2a34bf31df40cf07d96"
	dbname = "d4ehughfapgq0k"

)

var con *sql.DB

type Result struct {
	uid string
	type2 string
	quantity int
}

func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	var result Result
	resultSql:= "SELECT uid, type2, quantity FROM trnkelapabakar"
	
	err = db.QueryRow(resultSql).Scan(&result.uid, &result.type2, &result.quantity)
	
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}
	
	fmt.Println("Connected to postgres")
	con = db
	return db
	
}