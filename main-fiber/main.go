package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	// Import the generated docs package
	_ "github.com/pllus/main-fiber/docs"
)

// @title Go Fiber Swagger Example API
// @version 1.0
// @description This is a sample server for a Go Fiber API.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@example.com
// @host localhost:3000
// @BasePath /
func main() {
	// Your Fiber application setup
	app := fiber.New()

	// ... your routes

	// Add the Swagger UI route
	app.Get("/docs/*", swagger.HandlerDefault)

	// Start the server
	app.Listen(":3000")
}
