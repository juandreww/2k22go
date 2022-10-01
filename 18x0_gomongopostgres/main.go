package main

import (
	"fmt"
	"net/http"

	"github.com/juandreww/2k22go/18x0_gomongopostgres/prices"
	"github.com/juandreww/2k22go/18x0_gomongopostgres/staffs"
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

	http.HandleFunc("/staff", staffs.List)
	http.HandleFunc("/staff/new/form", staffs.NewStaffForm)
	http.HandleFunc("/staff/new/save", staffs.NewStaffSave)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I love food")
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}
