package main

import (
	"fmt"
	"net/http"
	// "encoding/json"
	"00_udemygointro/controller"
	"00_udemygointro/model"
	// "log"
	// "os"
)

func main() {
	mux := controller.RegisterApi()
	db := model.ConnectDB()
	defer db.Close()
	
	fmt.Println("Serving...");

	// for heroku
	// port := os.Getenv("PORT")
	// log.Fatal(http.ListenAndServe(":"+port, mux))

	// for Local
	http.ListenAndServe(":80", mux)
}