package validators

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ValidateBody struct{}

func (v *ValidateBody) ValidateBody(b IValidateBody) fiber.Handler {
	Validator := validator.New()

	return func(c *fiber.Ctx) error {

		body := b.MapType(c.Body())
		for _, b := range body {
			if err := Validator.Struct(b); err != nil {
				fmt.Println(err)
			}
		}
		return c.Next()
	}
}
