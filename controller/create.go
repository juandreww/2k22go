package controller

import (
	"net/http"
	// "encoding/json"
	"2k22go/model"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			if err := model.CreateKelapa(); err != nil {
				w.Write([]byte("Some error"))
				return
			}
		}
	}
}