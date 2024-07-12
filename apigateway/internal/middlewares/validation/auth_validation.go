package validation

import (
	"github.com/akselrivera/grpc-apigateway/internal/models"
	"github.com/gofiber/fiber/v2"
)

func ValidateLogin(c *fiber.Ctx) error {
	return validateBody(c, models.LoginRequest{})
}

func ValidateRegister(c *fiber.Ctx) error {
	return validateBody(c, models.RegisterRequest{})
}
