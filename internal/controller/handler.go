package controller

import (
	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *echo.Echo {
	app := echo.New()

	app.GET("/", h.Home)

	return app
}
