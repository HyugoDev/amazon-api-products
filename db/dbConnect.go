package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DbConnect() *mongo.Client {
	uri := "mongodb://localhost:27017/TestDatabaseAmazon?readPreference=primary&ssl=false&directConnection=true"

	Client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Ping the primary
	if err := Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged. (BASE DE DATOS MONGODB)")

	return Client
}

// func AddProduct(ctx context.Context, product string) (string, error) {
// 	productObj := map[string]interface{}{
// 		"ProductName": product,
// 	}

// 	collection := Client.Database("TestDatabaseAmazon").Collection("products")

// 	_, err := collection.InsertOne(ctx, productObj)

// 	if err != nil {
// 		return "", err
// 	}

// 	return product, nil
// }
