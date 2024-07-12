package validation

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type IError struct {
	Field string
	Tag   string
	Value string
}

func validateBody(c *fiber.Ctx, customStruct interface{}) error {
	var errors []*IError

	customType := reflect.TypeOf(customStruct)
	body := reflect.New(customType).Interface()
	c.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.Next()
}
