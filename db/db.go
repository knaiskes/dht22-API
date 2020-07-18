package db

import (
	"context"
	"fmt"
	//	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const (
	URI        = "mongodb://172.17.0.2:27017"
	TIMEOUT    = 10 * time.Second
	DATABASE   = "measurements"
	COLLECTION = "dht22"
)

var (
	collection *mongo.Collection
	ctx        = context.TODO()
)

func InitDB() {
	_, cancel := context.WithTimeout(ctx, TIMEOUT)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	/*
		defer func() {
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()
	*/

	err = client.Ping(ctx, readpref.Primary())
	defer cancel()

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to db")
	}

	collection = client.Database(DATABASE).Collection(COLLECTION)

	// Close connection
	/*
		databases, err := client.ListDatabaseNames(ctx, bson.M{})
		if err != nil {
			panic(err)
		}
		fmt.Println(databases)
	*/
}
