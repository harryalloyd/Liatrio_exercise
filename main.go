package main
import "github.com/gofiber/fiber/v2"
import "time"

func root(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "My name is Harrison Lloyd",
		"timestamp": time.Now().Unix(), //using unix to keep it as an int
	})
}

func main() {
	app := fiber.New()
	app.Get("/", root) // instance of a Fiber application
	app.Listen(":8080")
}
