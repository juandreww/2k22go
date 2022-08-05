package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"2k22go/views"
)

func main() {
	mux := http.NewServeMux()
	
	http.ListenAndServe(":80", mux)
}