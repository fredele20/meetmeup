package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type Database struct {
	client *mongo.Client
}

const DBNAME = "meetups"

var UserCollection *mongo.Collection
var MeetupsCollection *mongo.Collection

const USERCOLLECTION = "users"
const MEETUPCOLLECTION = "meetup"

//func init() {
//	//MONGODB := os.Getenv("MONGODB")
//	// set the client options
//	clientOptions := options.Client().ApplyURI(MONGODB)
//
//	client, err := mongo.Connect(context.TODO(), clientOptions)
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//
//	// check the connection
//	err = client.Ping(context.TODO(), nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("connected to the database successfully")
//	UserCollection = client.Database(DBNAME).Collection(USERCOLLECTION)
//	MeetupsCollection = client.Database(DBNAME).Collection(MEETUPCOLLECTION)
//}

func ConnectDB() *mongo.Collection {
	MONGODB := os.Getenv("MONGODB")
	// set the client options
	clientOptions := options.Client().ApplyURI(MONGODB)

	// connect to database
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection to database established")

	UserCollection = client.Database(DBNAME).Collection(USERCOLLECTION)
	MeetupsCollection = client.Database(DBNAME).Collection(MEETUPCOLLECTION)

	return UserCollection
	return MeetupsCollection
}
