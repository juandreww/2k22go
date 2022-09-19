package main

import (
	"log"
	"net/http"

	"rsc.io/letsencrypt"
)

func main() {
	http.HandleFunc("/", index)

	var enc letsencrypt.Manager
	if err := enc.CacheFile("letsencrypt.cache"); err != nil {
		log.Fatalln(err)
	}

	go http.ListenAndServe(":8080", http.HandlerFunc(letsencrypt.RedirectHTTP))

}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example of server.\n"))
}
