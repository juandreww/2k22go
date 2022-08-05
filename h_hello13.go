package main

import (
	// "fmt"
	"net/http"
	// "encoding/json"
	"2k22go/controller"
)

func main() {
	mux := controller.Register()
	
	http.ListenAndServe(":80", mux)
}