package validators

import (
	handlers "datalog-go/handler"
	"datalog-go/utils/logs"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateG1(c *fiber.Ctx) error {
	var Validator = validator.New()
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
	var Validator = validator.New()
	var errors []IError
	var body []handlers.Group2Dto

	c.BodyParser(&body)

	if len(body) > 0 {
		for _, b := range body {
			fmt.Println(b)
			err := Validator.Struct(b)
			if err != nil {
				for _, err := range err.(validator.ValidationErrors) {
					var el IError
					el.Field = err.Field()
					el.Tag = err.Tag()
					el.Value = err.Value()
					errors = append(errors, el)
					// logs.Error(el)
				}
				return c.Status(fiber.ErrBadRequest.Code).JSON(errors)
			}
		}
		return c.Next()
	}
	return c.SendStatus(fiber.ErrBadRequest.Code)
}
