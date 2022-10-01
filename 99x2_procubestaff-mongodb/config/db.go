package config

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var ctx = context.Background()

func init() {
	cOpt := options.Client().ApplyURI("mongodb://admin:123@localhost:27017")

	cl, err := mongo.NewClient(cOpt)
	checkError(err)

	err = cl.Connect(ctx)
	checkError(err)

	fmt.Println("Welcome to the mongodb.")
	DB = cl.Database("dbtest")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
