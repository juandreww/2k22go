package controller

import (
	"net/http"
	"encoding/json"
	"2k22go/model"
	"2k22go/views"
	// "log"
	// "github.com/davecgh/go-spew/spew"
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
			// uid := "568aa741-1063-4870-b3bb-0e3e81238c05"
			// uid := "bakar"
			// data, err := model.ReadSelected(uid)
			uid := r.URL.Query().Get("uid")
			data, err := model.ReadSelected(uid)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodDelete { 
			uid := r.URL.Query().Get("uid")
			if err := model.DeleteSelected(uid); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Status string `json:"Status"`
			}{"Item Deleted"})
		}
	}
}