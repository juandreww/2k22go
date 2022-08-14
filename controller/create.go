package controller

import (
	"net/http"
	"encoding/json"
	"2k22go/model"
	"2k22go/views"
	// "fmt"

)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.Kelapa{}
			json.NewDecoder(r.Body).Decode(&data)
			if err := model.CreateKelapa(data.type2, data.quantity); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}
	}
}