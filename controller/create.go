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
			log.Println(data)
			json.NewDecoder(r.Body).Decode(&data)
			log.Println(r.Body)
			log.Println(data)
			spew.Dump(r.Body)
			if err := model.CreateKelapa(data.Type2, data.Quantity); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}
	}
}