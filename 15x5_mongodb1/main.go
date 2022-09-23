package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	rt := httprouter.New()
	clientSession()

	// uc := controllers.NewUserController(getSession())
	// rt.GET("/user/:id", uc.GetUser)
	// rt.POST("/user", uc.CreateUser)
	// rt.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", rt)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27107")

	if err != nil {
		panic(err)
	}

	return s
}

func clientSession() {
	cOpt := options.Client().ApplyURI("mongodb://localhost:27017")

	cl, err := mongo.Connect(context.TODO(), cOpt)
	if err != nil {
		log.Fatal(err)
	}

	err = cl.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

}
