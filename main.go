package main

//import the func

import (
	// "github.com/HyugoDev/amazon-api-products/db"
	"github.com/HyugoDev/amazon-api-products/metodos"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	app.Get("/product", metodos.GetProduct)

	app.Post("product/add", metodos.PostProduct)

	app.Listen(":3000")
}
