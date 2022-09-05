package main

import (
	"html/template"
	"net/http"
	"fmt"
	"ioutil"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", lebegin)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func lebegin(w http.ResponseWriter, r *http.Request) {
	var s string
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
	}

	fmt.Println("\nfile:", f, "\nheader", h, "\nerr", err)
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s = string(bs)

	dst, err := os.Create(filepath.Join("./user/", h.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}