package example

import (
	"embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//go:embed test.json
var embeddedFiles embed.FS

// User represents the structure of a single user object in your JSON.
type User struct {
	ID        int      `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
}

// GetDataHandler reads the embedded JSON file and sends it as the response.
func GetDataHandler(c *fiber.Ctx) error {
	jsonData, err := embeddedFiles.ReadFile("test.json")

	if err != nil {
		log.Printf("Error reading embedded file: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Could not read data file.")
	}

	// Set the content type header and send the raw JSON data
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(http.StatusOK).Send(jsonData)
}

//go:embed post.json
var embed_file embed.FS

func GetDataHandler_post(c *fiber.Ctx) error {
	jsonData, err := embed_file.ReadFile("test.json")

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
