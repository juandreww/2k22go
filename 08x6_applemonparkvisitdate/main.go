package main

import (
	"net/http"
	"fmt"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/submitarrival", submitarrival)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Ready to serve..")
}

func submitarrival(w http.ResponseWriter, r *http.Request) {
	var s string
	if r.Method == "POST" {
		f, h, err := r.FormFile("file")
		HandleError(w, err)

		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

	}
}

func HandleError(w http.ResponseWriter,  err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}