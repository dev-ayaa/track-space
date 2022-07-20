package dbdriver

import (
	"context"
	"fmt"
	"log"
	"os"

	// "os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseConnection() *mongo.Client {
	mongodbUri := os.Getenv("MONGODBURI")
	if mongodbUri == "" {
		log.Println("mongodb cluster uri not found : ")
	}
	fmt.Println(mongodbUri)
	ctx, cancelCtx := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancelCtx()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://ayaaakinleye:2701Akin2000@cluster0.byrpjo8.mongodb.net/test"))

	if err != nil {

		log.Panic(err)
	}
	if err = client.Ping(ctx, nil); err != nil {
		log.Println("Failed to ping the database")
		panic(err)
	}
	log.Println("Database successfully pinged ! ")
	db, _ := client.ListDatabaseNames(ctx, bson.M{})
	fmt.Println(db)

	return client

}
