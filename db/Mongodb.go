package db

import (
	"context"
	"fmt"
	"log"

	"github.com/HyugoDev/amazon-api-products/ConfigApi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() *mongo.Client {
	Client, error := mongo.Connect(context.TODO(), options.Client().ApplyURI(ConfigApi.GetURLMongo()))

	if error != nil {
		log.Fatal(error)
		panic(error)
	} else {
		fmt.Println("Successfully connected (BASE DE DATOS MONGODB)")
		return Client
	}

}

func MongoDisconnection(client *mongo.Client) {
	error := client.Disconnect(context.TODO())

	fmt.Println("Successfully disconnected (BASE DE DATOS MONGODB)")

	if error != nil {
		panic(error)
	}
}
