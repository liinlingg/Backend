package example

import (
	"embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var embeddedFiles embed.FS

type DataResponse struct {
	Status string `json:"status" example:"success"`
	Data   []Item `json:"data"`
}

type Item struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"name" example:"Item A"`
}

func GetDataHandler(c *fiber.Ctx) error {
	jsonData, err := embeddedFiles.ReadFile("post.json")

	if err != nil {
		log.Printf("Error reading embedded file: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Could not read data file.")
	}

	// Set the content type header and send the raw JSON data
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(http.StatusOK).Send(jsonData)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/User", GetDataHandler)

	log.Fatal(app.Listen(":3000"))
}
