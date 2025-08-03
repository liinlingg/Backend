package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// A simple map to store valid client IDs.
// In a production environment, this would typically be fetched from a database.
var validClients = map[string]bool{
	"client_id_12345": true,
}

// AuthMiddleware checks for the presence and validity of a client ID.
func AuthMiddleware(c *fiber.Ctx) error {
	// Get the "X-Client-ID" header from the request.
	clientID := c.Get("X-Client-ID")

	// Check if the header is missing or if the client ID is not in our list of valid clients.
	if clientID == "" || !validClients[clientID] {
		// If the client ID is invalid or missing, return a 401 Unauthorized status.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid or missing client ID.",
		})
	}

	// If the client ID is valid, continue to the next middleware or route handler.
	return c.Next()
}