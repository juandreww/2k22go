package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/juandreww/2k22go/15x7_withoutmongodb/controllers"
	"github.com/julienschmidt/httprouter"
)

var ctx = context.Background()

func main() {
	rt := httprouter.New()

	fmt.Println("You are connected to MongoDB")
	uc := controllers.NewUserController()
	rt.GET("/user/:fname", uc.GetUser)
	// rt.POST("/user", uc.CreateUser)
	// rt.POST("/userupdate/:fname", uc.UpdateUser)
	// rt.DELETE("/user/:fname", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", rt)
}
