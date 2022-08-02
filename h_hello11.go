package main

import (
	"net/http"
	"fmt"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/getgoing", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received");
		fmt.Println(r.Method)
		w.Write([]byte("Hello World"))
	})
	fmt.Println("listening")
	http.ListenAndServe("localhost:80", mux)
}