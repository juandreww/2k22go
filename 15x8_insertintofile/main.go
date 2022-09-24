package main

import (
	"fmt"
	"net/http"

	"github.com/juandreww/2k22go/15x8_insertintofile/controllers"
	"github.com/juandreww/2k22go/15x8_insertintofile/models"
	"github.com/julienschmidt/httprouter"
)

// var ctx = context.Background()

func main() {
	rt := httprouter.New()

	fmt.Println("You are connected to MongoDB")
	uc := controllers.NewUserController(clientSession())
	rt.GET("/user/:fname", uc.GetUser)
	rt.POST("/user", uc.CreateUser)
	rt.POST("/file", uc.TextMe)
	rt.POST("/file2", uc.TextMe2)
	rt.DELETE("/user/:fname", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", rt)
}

func clientSession() map[string]models.Contacts {
	return make(map[string]models.Contacts)
}
