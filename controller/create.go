package controller

import (
	"net/http"
	"encoding/json"
	"2k22go/model"
	"2k22go/views"
	"log"
	"github.com/davecgh/go-spew/spew"
	// "fmt"

)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.Kelapa{}
			json.NewDecoder(r.Body).Decode(&data)
			if err := model.CreateKelapa(data.Type2, data.Quantity); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			data, err := model.ReadAll()
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}