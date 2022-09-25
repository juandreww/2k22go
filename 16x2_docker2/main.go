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

	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	mux.HandleFunc("/index1", index)
	mux.HandleFunc("/index2", index2)
	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "http method not allowed", http.StatusBadRequest)
			return
		}

		text := "Hello world"
		if instanceID != "" {
			text = text + ". From " + instanceID
		}

		w.Write([]byte(text))
	})

	server := new(http.Server)
	server.Handler = mux
	server.Addr = "0.0.0.0" + port
	log.Println("Server starting at ", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, _ *http.Request) {
	hello := "Hello from Docker 1"

	w.Write([]byte(hello))
}

func index2(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello from docker index 2")
}
