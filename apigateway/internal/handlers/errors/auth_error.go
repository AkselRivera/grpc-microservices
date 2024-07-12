package errors

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func HandleAuthError(c *fiber.Ctx, err error) error {
	log.Errorf("Error getting auth: %v", err)

	if strings.Contains(err.Error(), "invalid token") {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token, please login",
		})
	}

	if strings.Contains(err.Error(), "user already exists") {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "User already exists, please try with another email",
		})
	}

	if strings.Contains(err.Error(), "error hashing password") {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	if strings.Contains(err.Error(), "invalid credentials") {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"message": "Couldn't connect with the auth - microservice",
	})
}
