package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.ListenAndServe(":80", nil)
}