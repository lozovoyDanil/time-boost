package controller

import (
	"fmt"
	"log/slog"
	"net/http"
	"time_boost/internal/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Register(c *fiber.Ctx) error {
	var u model.User

	if err := c.BodyParser(&u); err != nil {
		slog.Error("Registration", err)

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	c.Status(http.StatusOK)

	return nil
}

type loginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *Handler) Login(c *fiber.Ctx) error {
	input := new(loginInput)

	if err := c.BodyParser(input); err != nil {
		slog.Error("Login", err)

		return c.Status(http.StatusBadRequest).Render("login", err)
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		slog.Error("Login", err)

		return c.Status(http.StatusBadRequest).Render("login", err)
	}

	slog.Info("Login", "SUCCESS", fmt.Sprintf("%+v", input))
	return c.Status(http.StatusOK).Render("login", "SUCCESS")
}
