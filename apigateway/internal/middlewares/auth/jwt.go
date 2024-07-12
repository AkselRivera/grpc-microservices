package auth

import (
	"net/http"

	"github.com/akselrivera/grpc-apigateway/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoute(c *fiber.Ctx) error {

	if handlers.ValidateTokenMiddleware(c) {
		return c.Next()
	}

	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
		"message": "Unauthorized",
	})
}
