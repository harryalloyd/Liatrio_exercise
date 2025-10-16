package main
import "github.com/gofiber/fiber/v2"
import "time"
import "encoding/json"

type Payload struct {
	Message string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func root(c *fiber.Ctx) error {
	p := Payload{
		Message:   "My name is Harrison Lloyd",
		Timestamp: time.Now().UnixMilli(), // milliseconds since epoch
	}
	b, _ := json.Marshal(p) // Marshal is minified by default (no spaces/newlines)
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(b)
}

func main() {
	app := fiber.New()
	app.Get("/", root) // instance of a Fiber application
	app.Listen(":8080")
}
