package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/juandreww/2k22go/15x6_mongodb2/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

func main() {
	rt := httprouter.New()
	db, err := clientSession()
	fmt.Println(db)
	if err != nil {
		panic(err)
	}

	uc := controllers.NewUserController(db)
	rt.GET("/user/:id", uc.GetUser)
	rt.POST("/user", uc.CreateUser)
	rt.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", rt)
}

func clientSession() (*mongo.Database, error) {
	cOpt := options.Client().ApplyURI("mongodb://admin:123@localhost:27017")

	cl, err := mongo.NewClient(cOpt)
	if err != nil {
		return nil, err
	}

	err = cl.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return cl.Database("dbtest"), nil
}
