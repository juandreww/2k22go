package main

import (
	"net/http"
	"encoding/json"
	"2k22go/structs"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code : http.StatusOK,
				Body : "pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe(":80", nil)
}