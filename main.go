package main
import "github.com/gofiber/fiber/v2"
import "time"
import "encoding/json"

type Payload struct {
	Message string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func root(c *fiber.Ctx) error {
	resp := Payload{
		Message: "My name is Harrison Lloyd",
		Timestamp: time.Now().UnixMilli(), 
	}
	b, _ := json.Marshal(resp) 
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(b)
}

func main() {
	app := fiber.New()
	app.Get("/", root) 
	app.Listen(":8080")
}
