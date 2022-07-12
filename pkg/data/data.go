package data

import (
	"go.mongodb.org/mongo-driver/mongo"
)

//Setting up the database for the user data collection
func UserData(dbClient *mongo.Client, collectionName string) *mongo.Collection {
	var userCollection *mongo.Collection
	userCollection = dbClient.Database("Track_Space").Collection(collectionName)
	return userCollection
}

//Setting up the database for the mail data collection
func MailData(dbClient *mongo.Client, collectionName string) *mongo.Collection {
	var mailCollection *mongo.Collection
	mailCollection = dbClient.Database("Track_Space").Collection(collectionName)
	return mailCollection
}
