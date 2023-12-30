package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Register(c *fiber.Ctx) error {
	err := c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successful registration",
	})

	return err
}
