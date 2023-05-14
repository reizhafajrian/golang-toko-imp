package provider

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectingMongo() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://reizha77:RwGBmJRmnLliT9lr@cluster0.guv2o.mongodb.net/?retryWrites=true&w=majority")

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
