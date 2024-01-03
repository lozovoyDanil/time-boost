package controller

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// render is a shortcut for rendering a component.
func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
