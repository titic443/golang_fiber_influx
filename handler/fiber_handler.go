package handlers

import (
	"datalog-go/services"
	"datalog-go/utils/logs"
	"datalog-go/utils/validators"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	g1 services.IGroup1
	g2 services.IGroup2
	g3 services.IGroup3
	g4 services.IGroup4
	g5 services.IGroup5
}

func NewHandler(g1 services.IGroup1, g2 services.IGroup2, g3 services.IGroup3, g4 services.IGroup4, g5 services.IGroup5) handler {
	return handler{
		g1: g1,
		g2: g2,
		g3: g3,
		g4: g4,
		g5: g5,
	}
}

func (h handler) InsertG1(c *fiber.Ctx) error {
	validateBody := c.Locals("validateBody").([]interface{})
	body := []validators.Group1Dto{}
	c.BodyParser(&body)

	for _, b := range validateBody {
		err := h.g1.InsertDataToAmita(b)
		if err != nil {
			logs.Error(err)
			return c.Status(fiber.ErrInternalServerError.Code).JSON(err)
		}
	}
	return nil
}

func (h handler) InsertG2(c *fiber.Ctx) error {
	validateBody := c.Locals("validateBody").([]interface{})
	body := []validators.Group2Dto{}
	err := c.BodyParser(&body)
	if err != nil {
		logs.Error(err)
	}
	for _, b := range validateBody {
		err := h.g2.InsertDataToAmita(b)
		if err != nil {
			logs.Error(err)
			return c.Status(fiber.ErrInternalServerError.Code).JSON(err)
		}
	}
	return nil
}

func (h handler) InsertG3(c *fiber.Ctx) error {
	validateBody := c.Locals("validateBody").([]interface{})
	body := []validators.Group3Dto{}
	c.BodyParser(&body)

	for _, b := range validateBody {
		err := h.g3.InsertDataToAmita(b)
		if err != nil {
			logs.Error(err)
			return c.Status(fiber.ErrInternalServerError.Code).JSON(err)
		}
	}
	return nil
}

func (h handler) InsertG4(c *fiber.Ctx) error {
	validateBody := c.Locals("validateBody").([]interface{})
	body := []validators.Group4Dto{}
	c.BodyParser(&body)

	for _, b := range validateBody {
		err := h.g4.InsertDataToAmita(b)
		if err != nil {
			logs.Error(err)
			return c.Status(fiber.ErrInternalServerError.Code).JSON(err)
		}
	}
	return nil
}

func (h handler) InsertG5(c *fiber.Ctx) error {
	validateBody := c.Locals("validateBody").([]interface{})
	body := []validators.Group5Dto{}
	c.BodyParser(&body)

	for _, b := range validateBody {
		err := h.g5.InsertDataToAmita(b)
		if err != nil {
			logs.Error(err)
			return c.Status(fiber.ErrInternalServerError.Code).JSON(err)
		}
	}
	return nil
}
