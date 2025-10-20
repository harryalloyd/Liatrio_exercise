package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func root(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "My name is Harrison Lloyd",
		"timestamp": time.Now().UnixMilli(),
		"deployed_at": time.Now().Format(time.RFC3339),
	})
}

func main() {
	app := fiber.New()
	app.Get("/", root) // instance of a Fiber application
	app.Listen(":8080")
}