package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	rt := httprouter.New()
	uc := controllers.NewUserController()
	rt.GET("/user/:id", uc.GetUser)
	rt.POST("/user", uc.CreateUser)
	rt.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", rt)
}
