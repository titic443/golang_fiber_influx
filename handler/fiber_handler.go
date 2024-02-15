package handlers

import (
	"datalog-go/services"
	"datalog-go/utils/logs"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	g1 services.IGroup1
	g2 services.IGroup2
}

func NewHandler(g1 services.IGroup1, g2 services.IGroup2) handler {
	return handler{
		g1: g1,
		g2: g2,
	}
}

func (h handler) InsertG1(c *fiber.Ctx) error {
	body := []Group1Dto{}
	c.BodyParser(&body)

	for _, b := range body {
		err := h.g1.InsertDataToAmita(b)
		if err != nil {
			logs.Error(err)
			return c.Status(fiber.ErrInternalServerError.Code).JSON(err)
		}
	}
	return nil
}
