package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/juandreww/2k22go/15x6_mongodb2/controllers"
	"github.com/juandreww/2k22go/15x7_withoutmongodb/models"
	"github.com/julienschmidt/httprouter"
)

var ctx = context.Background()

func main() {
	rt := httprouter.New()
	db := clientSession()

	fmt.Println("You are connected to MongoDB")
	uc := controllers.NewUserController(db)
	rt.GET("/user/:fname", uc.GetUser)
	rt.POST("/user", uc.CreateUser)
	rt.POST("/userupdate/:fname", uc.UpdateUser)
	rt.DELETE("/user/:fname", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", rt)
}

func clientSession() map[string]models.Contacts {
	return make(map[string]models.Contacts)
}
