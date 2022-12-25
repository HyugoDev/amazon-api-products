package metodos

import (
	"context"

	"github.com/HyugoDev/amazon-api-products/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func PostProduct(p *fiber.Ctx, dbClient *mongo.Client) error {
	var product models.Product

	p.BodyParser(&product)

	coll := dbClient.Database("TestDatabaseAmazon").Collection("products")
	result, err := coll.InsertOne(context.TODO(), product)

	if err != nil {
		panic(err)
	}

	return p.JSON(&fiber.Map{
		"status": result.InsertedID,
		"data":   product,
	})
}
