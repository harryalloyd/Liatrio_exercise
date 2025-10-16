package main
import "github.com/gofiber/fiber/v2"
import "time"

type Payload struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func root(c *fiber.Ctx) error {
	return c.JSON(Payload{
		Message: "My name is Harrison Lloyd",
		Timestamp: time.Now().UnixMilli(), 
	})
}

func main() {
	app := fiber.New()
	app.Get("/", root) // instance of a Fiber application
	app.Listen(":8080")
}
