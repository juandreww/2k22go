package main

import (
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
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
		if f == nil {
			http.Error(w, "Sorry but file is nil", http.StatusInternalServerError)
			return
		}
		HandleError(w, err)

		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)
		str, err := ioutil.ReadAll(f)
		HandleError(w, err)
		// fmt.Println(str)

		s = string(str)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<h2>Please submit your arrival datetime at our Lemon Park ! </h2>
		<form method="POST" enctype="multipart/form-data">
		<input type="file" name="file">
		<input type="submit">
		</form>
	<br>`+s)
}

func HandleError(w http.ResponseWriter,  err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}