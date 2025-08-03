package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/pllus/main-fiber/docs"
	"github.com/pllus/main-fiber/middleware"
	"github.com/pllus/main-fiber/routes/example"
)

func main() {
	app := fiber.New()

	app.Get("/docs/*", swagger.HandlerDefault)

	app.Get("/hello", getHello)
	privateRoutes := app.Group("/", middleware.AuthMiddleware)
	app.Get("/User", example.GetDataHandler)
	privateRoutes.Get("/Protected", example.GetDataHandler)

	// app.Get("/Post",)

	//server variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}

func getHello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
