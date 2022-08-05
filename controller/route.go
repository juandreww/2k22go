package controller

import (
	"net/http"
	"encoding/json"
	"2k22go/views"
	"fmt"
)

func Startup() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Println("Request received")
			data := views.Response{
				Code : http.StatusOK,
				Body : "pong",
			}
			
			json.NewEncoder(w).Encode(data)
		}
	})
}