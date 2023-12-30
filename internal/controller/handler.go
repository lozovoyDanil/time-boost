package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *fiber.App {
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	auth := app.Group("/auth")
	{
		auth.Post("/register", h.Register)
		auth.Post("/login", h.Login)
	}

	return app
}
