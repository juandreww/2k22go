package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/juandreww/2k22go/15x3_postdelete1/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	rt := httprouter.New()
	rt.GET("/index", index)
	rt.GET("/user/:id", getUser)
	rt.POST("/user", createUser)
	rt.DELETE("/user/:id", deleteUser)
	http.ListenAndServe("localhost:8080", rt)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	str := `<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Index</title>
		</head>
		<body>
		<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
		</body>
		</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(str))
}

func getUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.NewUser{
		Name:   "Avaleska Zhang",
		Gender: "female",
		Age:    18,
		Id:     p.ByName("id"),
	}

	js, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", js)
}

func createUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(req.Body).Decode(&u)

	u.Id = "007"
	js, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", js)
}

func deleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Write code to delete user\n")
}
