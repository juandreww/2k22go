package controller

import (
	"net/http"
	"encoding/json"
	"2k22go/views"
	"fmt"
)

func ping() http.HandlerFunc {
	fmt.Println("ping")
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Println("Request received")
			data := views.Response{
				Code : http.StatusOK,
				Body : "pong",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}