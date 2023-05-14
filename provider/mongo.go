package provider

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectingMongo() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Error pinging MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}
