package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	mux.HandleFunc("/authenticate", auth)
	http.ListenAndServe(":8080", nil)
}

func getCode(str string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
