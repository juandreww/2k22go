package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	rt := httprouter.New()
	rt.GET("/index", index)
	rt.GET("/user/:id", getUser)
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
