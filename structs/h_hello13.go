package main

// import (
// 	"fmt"
// 	"net/http"
// 	// "encoding/json"
// 	"2k22go/controller"
// 	"2k22go/model"
// )

// func main() {
// 	mux := controller.Register()
// 	db := model.Connect()
// 	defer db.Close()
// 	fmt.Println("Serving...");
// 	http.ListenAndServe(":80", mux)
// }