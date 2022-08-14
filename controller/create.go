package controller

import (
	"net/http"
	"encoding/json"
	// "2k22go/model"
	"2k22go/views"
	"fmt"

)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Println("Request received")
			data := views.Response{
				Code : http.StatusOK,
				Body : "pong",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} else {
			fmt.Println("Request different")
			data := views.Response{
				Code : http.StatusOK,
				Body : "pong different",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}

	// return func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == http.MethodGet {
	// 		if err := model.CreateKelapa(); err != nil {
	// 			w.Write([]byte("Some error"))
	// 			return
	// 		}
	// 	}
	// }
}