package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	//create instance of fiber
	app := fiber.New()

	//create http handlers
	app.Get("/testApi", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Go fiber first api created",
		})
	})

	//listen on port
	app.Listen(":3000")

}
