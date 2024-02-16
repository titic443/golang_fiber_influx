package validators

import (
	"datalog-go/utils/logs"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ValidateBody struct{}

func (v *ValidateBody) ValidateBody(b IValidateBody) fiber.Handler {
	Validator := validator.New()

	return func(c *fiber.Ctx) error {

		body := b.MapType(c.Body())
		if body == nil {
			return c.SendStatus(fiber.ErrBadRequest.Code)
		}
		for _, b := range body {
			if err := Validator.Struct(b); err != nil {
				logs.Error(err)
			}
		}
		logs.Info("Validate success")
		return c.Next()
	}
}
