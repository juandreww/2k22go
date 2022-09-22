package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	rt := httprouter.New()
	rt.GET("/index", index)
	http.ListenAndServe("localhost:8080", rt)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!")
}
