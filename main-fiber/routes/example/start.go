package example

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

// DataResponse is a model that describes the structure of the JSON data.
// It's used for documentation purposes by Swagger.
type DataResponse struct {
	Status string `json:"status" example:"success"`
	Data   []Item `json:"data"`
}

// Item is a model for an item in the data array.
type Item struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"name" example:"Item A"`
}

// readJSONFile is a helper function that reads a file and returns its content.
func readJSONFile(filePath string) ([]byte, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// @Summary Get static data
// @Description Reads and sends the content of a static data.json file.
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} DataResponse "The data from the JSON file"
// @Failure 500 {string} string "Could not read data file."
// @Router /User [get]
func GetDataHandler(c *fiber.Ctx) error {
	jsonData, err := readJSONFile("/mock_data/post.json")
	if err != nil {
		log.Printf("Error reading data.json: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Could not read data file.")
	}

	// Set the content type header and send the raw JSON data
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Send(jsonData)
}
