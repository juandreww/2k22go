package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"github.com/google/uuid"
)

const (
	host = "ec2-52-73-155-171.compute-1.amazonaws.com"
	port = 5432
	user = "'awtpbzyctlpydm'"
	password = "a7bf40c39496f73a03e7412befbc787d29138445d7fce2a34bf31df40cf07d96"
	dbname = "d4ehughfapgq0k"
)

type Kelapa struct {
	type2 string
	quantity float64
}

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
	
	var kelapa Kelapa
	kelapaSql := "SELECT type2, quantity FROM trnkelapabakar"

	err = db.QueryRow(kelapaSql).Scan(&kelapa.type2, &kelapa.quantity)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	fmt.Printf("Hi %s welcome to my channel, quantity is %.2f\n", kelapa.type2, kelapa.quantity)

	// $query := "INSERT INTO trnkelapabakar ("
	id := uuid.New()
	fmt.Println(id.String())
}