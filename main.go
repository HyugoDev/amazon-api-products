package main

import (
	"github.com/HyugoDev/amazon-api-products/ConfigApi"
	"github.com/HyugoDev/amazon-api-products/db"
	"github.com/HyugoDev/amazon-api-products/metodos"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	//Cargo el archivo .env
	ConfigApi.LoadEnv()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	//mostrar en consola las llamadas que se han echo a las rutas
	app.Use(logger.New(logger.Config{
		Format:     "${latency}  Fecha: ${time} - IP: ${ip} - ${method} ${path} ${status} \n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/CostaRica",
	}))

	// base data
	dbClient := db.MongoConnection()
	defer db.MongoDisconnection(dbClient)

	//Rutas
	app.Get("/product", func(c *fiber.Ctx) error {
		return metodos.GetProduct(c, dbClient)
	})

	app.Post("/product/add", func(c *fiber.Ctx) error {
		return metodos.PostProduct(c, dbClient)
	})

	//Escucha el puerto
	app.Listen(ConfigApi.GetPORT())
}
