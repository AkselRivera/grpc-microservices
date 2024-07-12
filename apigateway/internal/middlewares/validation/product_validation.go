package validation

import (
	"github.com/akselrivera/grpc-apigateway/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

func ValidateProduct(c *fiber.Ctx) error {
	return validateBody(c, models.ProductComplete{})

}

func ValidateProductWithoutId(c *fiber.Ctx) error {
	return validateBody(c, models.ProductWithoutID{})
}

func ValidateProductStock(c *fiber.Ctx) error {
	return validateBody(c, models.ProductStock{})
}

// func ValidateProductWithoutId(c *fiber.Ctx) error {
// 	var errors []*IError
// 	body := new(models.ProductWithoutID)
// 	c.BodyParser(&body)

// 	err := Validator.Struct(body)
// 	if err != nil {
// 		for _, err := range err.(validator.ValidationErrors) {
// 			var el IError
// 			el.Field = err.Field()
// 			el.Tag = err.Tag()
// 			el.Value = err.Param()
// 			errors = append(errors, &el)
// 		}
// 		return c.Status(fiber.StatusBadRequest).JSON(errors)
// 	}
// 	return c.Next()
// }
