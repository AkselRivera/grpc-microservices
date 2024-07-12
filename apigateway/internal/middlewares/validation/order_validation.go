package validation

import (
	"github.com/akselrivera/grpc-apigateway/internal/models"
	"github.com/gofiber/fiber/v2"
)

func ValidateOrder(c *fiber.Ctx) error {
	return validateBody(c, models.Order{})
}

func ValidateChangeStatus(c *fiber.Ctx) error {
	return validateBody(c, models.ChangeOrderStatus{})
}
