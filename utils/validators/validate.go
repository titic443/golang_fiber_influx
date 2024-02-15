package validators

import (
	handlers "datalog-go/handler"
	"datalog-go/utils/logs"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

func ValidateG1(c *fiber.Ctx) error {
	var errors []IError
	body := []handlers.Group1Dto{}
	c.BodyParser(&body)

	if len(body) > 0 {
		for _, b := range body {

			err := Validator.Struct(b)
			if err != nil {
				for _, err := range err.(validator.ValidationErrors) {
					var el IError
					el.Field = err.Field()
					el.Tag = err.Tag()
					el.Value = err.Value()
					errors = append(errors, el)
					logs.Error(el)
				}
				return c.Status(fiber.ErrBadRequest.Code).JSON(errors)
			}
		}
		return c.Next()
	}
	return c.SendStatus(fiber.ErrBadRequest.Code)
}

func ValidateG2(c *fiber.Ctx) error {
	var errors []IError
	body := []handlers.Group2Dto{}
	c.BodyParser(&body)

	if len(body) > 0 {
		for _, b := range body {

			err := Validator.Struct(b)
			if err != nil {
				for _, err := range err.(validator.ValidationErrors) {
					var el IError
					el.Field = err.Field()
					el.Tag = err.Tag()
					el.Value = err.Value()
					errors = append(errors, el)
					logs.Error(el)
				}
				return c.Status(fiber.ErrBadRequest.Code).JSON(errors)
			}
		}
		return c.Next()
	}
	return c.SendStatus(fiber.ErrBadRequest.Code)
}
