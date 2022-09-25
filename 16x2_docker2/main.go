package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env is required")
	}

	instanceID := os.Getenv("INSTANCE_ID")

	http.HandleFunc("/", index)
	http.HandleFunc("/index1", index)
	http.HandleFunc("/index2", index2)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello from docker index 1")
}

func index2(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello from docker index 2")
}
