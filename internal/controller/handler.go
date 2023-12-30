package controller

import "github.com/gofiber/fiber/v2"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *fiber.App {
	app := fiber.New()

	auth := app.Group("/auth")
	{
		auth.Post("/register", h.Register)
		// auth.Post("/login", h.Login)
	}

	return app
}
