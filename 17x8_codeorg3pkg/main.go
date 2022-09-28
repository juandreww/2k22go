package main

import (
	"net/http"

	"github.com/juandreww/2k22go/17x8_codeorg3pkg/prices"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/index", prices.Index)
	http.HandleFunc("/index/show", prices.IndexShow)
	http.HandleFunc("/index/create", prices.IndexCreateForm)
	http.HandleFunc("/index/create/process", prices.IndexCreateProcess)
	http.HandleFunc("/index/update", prices.IndexUpdateForm)
	http.HandleFunc("/index/update/process", prices.IndexUpdateProcess)
	http.HandleFunc("/index/delete/process", prices.IndexDeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}
