package metodos

import (
	"github.com/gofiber/fiber/v2"
)

func PostProduct(p *fiber.Ctx) error {

	return p.SendString("Hello, World!")
}
