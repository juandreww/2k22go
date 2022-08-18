package main

import (
	"fmt"
	"net/http"
	// "encoding/json"
	"2k22go/controller"
	"2k22go/model"
	"log"
	"os"
)

func main() {
	mux := controller.RegisterApi()
	db := model.ConnectDB()
	defer db.Close()
	
	fmt.Println("Serving...");
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, mux))
	// http.ListenAndServe(":80", mux)
}