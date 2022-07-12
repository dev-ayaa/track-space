package dbdriver

import (
	"context"
	"fmt"
	"log"
	// "os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var Client *mongo.Client = DatabaseConnection()

func DatabaseConnection() *mongo.Client {
	
	// mongodb_uri := os.Getenv("MONGODB_URI")
	// if mongodb_uri == "" {
	// 	log.Println("mongodb cluster uri not found : ")
	// }

	// authMechanism := options.Credential{
	// 	AuthMechanism: "SCRAM-SHA-256",
	// 	AuthSource:    "Users",
	// 	Password:      "2701Akin2000",
	// 	Username:      "ayaaakinleye",
	// }

	ctx, cancelCtx := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancelCtx()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://ayaaakinleye:2701Akin2000@cluster0.byrpjo8.mongodb.net/test"))
	//.SetAuth(authMechanism))
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

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
			return
		}
	}()

	return client

}
