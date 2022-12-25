package metodos

import (
	"context"

	"github.com/HyugoDev/amazon-api-products/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetProduct muestra los product
func GetProduct(p *fiber.Ctx, dbclient *mongo.Client) error {
	var products []models.Product

	coll := dbclient.Database("TestDatabaseAmazon").Collection("products")

	results, err := coll.Find(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}

	for results.Next(context.TODO()) {
		var product models.Product
		results.Decode(&product)
		products = append(products, product)
	}

	return p.Status(fiber.StatusOK).JSON(&fiber.Map{
		"Products": products,
	})
}
