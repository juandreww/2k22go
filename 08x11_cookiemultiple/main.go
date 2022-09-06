package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", setmultiple)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func setmultiple(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-name",
		Value: "David McGregor",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "my-order",
		Value: "Avalon Lemon ",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "my-quantity",
		Value: "1",
	})

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}