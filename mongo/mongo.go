package mongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectionString = "mongodb://staging:27017"
	dbName           = "netflix"
	collectionName   = "watchlist"
)

// MOST IMPORTANT
var (
	collection *mongo.Collection
)

// connect with mongo
func init() {
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connectToMongo
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo connection success")
	collection = client.Database(dbName).Collection(collectionName)

	// collection instance
	fmt.Println("Collection instance is ready")
}
