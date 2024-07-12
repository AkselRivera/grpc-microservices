package errors

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func HandleProductError(c *fiber.Ctx, err error) error {
	log.Errorf("Error getting product: %v", err)

	if strings.Contains(err.Error(), "Product not found") {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	if strings.Contains(err.Error(), "Product is inactive") {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Product is inactive, please use another product",
		})
	}

	if strings.Contains(err.Error(), "Not enough stock") {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "Product out of stock, please use another product",
		})
	}

	if strings.Contains(err.Error(), "Invalid action") {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "Unkown action to perform, try again",
		})
	}

	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"message": "Couldn't connect with the product - microservice",
	})
}
