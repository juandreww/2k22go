package controller

import (
	"net/http"
	// "encoding/json"
	// "2k22go/views"
	// "fmt"
)

func RegisterApi() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/", create())
	return mux
}