package databases

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var Collection *mongo.Collection

func ConnectMongo() {

	connectionURI := os.Getenv("MONGO_CONNECTION_URL")
	clientOptions := options.Client().ApplyURI(connectionURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to MongoDB!")

	// Access a database and collection
	Collection = client.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_COLLECTION"))

}
