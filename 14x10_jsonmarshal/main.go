package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encode)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>INDEX</title>
	</head>
	<body>
	You are at index
	</body>
	</html>`
	w.Write([]byte(s))
}

func marshal(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	p1 := person{
		Fname: "Louis",
		Lname: "Vuitton",
		Items: []string{"Modern", "Bag", "Very original"},
	}
	js, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(js)
}

func encode(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "Louis",
		Lname: "Vuitton",
		Items: []string{"Modern", "Bag", "Very Original"},
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
