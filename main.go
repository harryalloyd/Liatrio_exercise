package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

// Handles a web request and sends back a JSON object to the caller. 
// It builds that object using a map literal (key–value dict) containing the message
func showInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{ 
		"message": "My name is Harrison Lloyd",
		"timestamp": time.Now().UnixMilli(), //Returns the current time in milliseconds
		"deployed_at": time.Now().Format(time.RFC3339), //Returns the year–month–day and hours-minutes-seconds 
	})
}

func main() {
	app:=fiber.New() // Creates a new Fiber web application
	app.Get("/", showInfo) // Registers the showInfo handler for HTTP GET requests on "/"
	app.Listen(":8080") // Starts the HTTP server on port 8080
}