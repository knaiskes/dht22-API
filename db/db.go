package db

import (
	"context"
	"fmt"
	"github.com/KNaiskes/measurementsTH-API/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const (
	URI         = "mongodb://172.17.0.2:27017"
	TIMEOUT     = 10 * time.Second
	GET_TIMEOUT = 30 * time.Second
	DATABASE    = "measurements"
	COLLECTION  = "dht22"
)

var (
	collection  *mongo.Collection
	ctx         = context.TODO()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(URI))
	_, cancel   = context.WithTimeout(ctx, TIMEOUT)
)

func Connect() {
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	defer cancel()

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to db")
	}
}

func GetAll() []models.Measurement {
	collection = client.Database(DATABASE).Collection(COLLECTION)
	fmt.Printf("%T\n", collection)
	results := []models.Measurement{}

	ctx, cancel = context.WithTimeout(context.Background(), GET_TIMEOUT)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println(err)
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		row := models.Measurement{}

		err = cur.Decode(&row)
		if err != nil {
			fmt.Println(err)
		}
		results = append(results, row)
	}

	return results
}
