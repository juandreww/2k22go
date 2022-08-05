package controller

import (
	"net/http"
	// "encoding/json"
	// "2k22go/views"
	// "fmt"
)

func Startup() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	return mux
}