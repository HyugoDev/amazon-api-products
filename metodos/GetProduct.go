package metodos

import (
	"github.com/HyugoDev/amazon-api-products/models"
	"github.com/gofiber/fiber/v2"
)

// GetProduct muestra los product
func GetProduct(p *fiber.Ctx) error {
	var products []models.Product

	// Client := db.DbConnect()

	// coll := Client.Collection("products")

	// results, err := coll.Find(context.TODO(), bson.M{})

	// if err != nil {
	// 	panic(err)
	// }

	// for results.Next(context.TODO()) {
	// 	var product models.Product
	// 	results.Decode(&product)
	// 	products = append(products, product)
	// }

	return p.Status(fiber.StatusOK).JSON(&fiber.Map{
		"Products": products,
	})
}
