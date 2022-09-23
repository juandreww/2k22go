package main

import (
	"net/http"

	"github.com/juandreww/2k22go/15x5_mongodb1/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	rt := httprouter.New()
	uc := controllers.NewUserController(getSession())
	rt.GET("/user/:id", uc.GetUser)
	rt.POST("/user", uc.CreateUser)
	rt.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", rt)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return s
}
