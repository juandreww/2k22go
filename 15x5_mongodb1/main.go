package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/juandreww/2k22go/15x5_mongodb1/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cred options.Credential

cred.

type Contacts struct {
	Fname string `bson:"Fname"`
	Lname string `bson:"Lname"`
}

func main() {
	rt := httprouter.New()
	db, err := clientSession()
	if err != nil {
		panic(err)
	}

	uc := controllers.NewUserController(db)
	rt.GET("/user/:id", uc.GetUser)
	rt.POST("/user", uc.CreateUser)
	rt.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", rt)
}

func clientSession() (*mongo.Collection, error) {
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
	collection := cl.Database("dbtest").Collection("contacts")

	_, err = cl.Database("dbtest").Collection("contacts").InsertOne(context.TODO(), Contacts{"Abigail", "Ayu"})
	if err != nil {
		log.Fatal(err)
	}

	return collection, nil
}
