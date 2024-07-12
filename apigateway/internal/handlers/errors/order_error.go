package errors

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func HandleOrderError(c *fiber.Ctx, err error) error {
	log.Errorf("Error getting order: %v", err)

	if strings.Contains(err.Error(), "product not found") {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	if strings.Contains(err.Error(), "order not found") {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Order not found",
		})
	}

	if strings.Contains(err.Error(), "product is inactive") {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Product is inactive, please use another product",
		})
	}

	if strings.Contains(err.Error(), "product out of stock") {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "Product out of stock, please use another product",
		})
	}

	if strings.Contains(err.Error(), "order cannot be cancelled") {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "Please use cancel order route instead",
		})
	}

	if strings.Contains(err.Error(), "order already cancelled") {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "Ups! Order already cancelled",
		})
	}

	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"message": "Couldn't connect with the order - microservice",
	})
}
