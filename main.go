package main

import (
	"fmt"
	"net/http"
	// "encoding/json"
	"2k22go/controller"
	"2k22go/model"
)

func main() {
	mux := controller.RegisterApi()
	db := model.ConnectDB()
	defer db.Close()
	
	fmt.Println("Serving...");
	http.ListenAndServe(":80", mux)
}