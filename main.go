package main
import "github.com/gofiber/fiber/v2"
import "time"

type Payload struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		p := Payload{
			Message:   "My name is Harrison Lloyd",
			Timestamp: time.Now().Unix(), 
		}
		return c.JSON(p) 
	})

	app.Listen(":8080")
}
