package controller

import (
	"net/http"
	// "encoding/json"
	"2k22go/model"
	"fmt"
)

func create() http.HandlerFunc {
	fmt.Println("create")
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if err := model.CreateKelapa(); err != nil {
				w.Write([]byte("Some error"))
				return
			} else {
				w.Write([]byte("No error"))
			}
		}
	}
}