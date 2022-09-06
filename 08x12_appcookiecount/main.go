package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
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
		Value: "Avalon Lemon",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "my-quantity",
		Value: "1",
	})

	c1, err := req.Cookie("my-logincount")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "You have logged in ", c1.Value, " times")
	}

	count, err := strconv.Atoi(c1.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	c1.Value = strconv.Itoa(count)

	http.SetCookie(w, &http.Cookie{
		Name:  "my-logincount",
		Value: c1.Value,
	})

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("my-name")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #1:", c1)
	}

	c2, err := req.Cookie("my-order")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #2:", c2)
		fmt.Println(c2.Value)
	}

	c3, err := req.Cookie("my-quantity")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #3:", c3)
	}
}