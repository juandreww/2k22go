package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	host = "ec2-52-73-155-171.compute-1.amazonaws.com"
	port = 5432
	user = "'awtpbzyctlpydm'"
	password = "a7bf40c39496f73a03e7412befbc787d29138445d7fce2a34bf31df40cf07d96"
	dbname = "d4ehughfapgq0k"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=require",
    host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected successfully")
	
}