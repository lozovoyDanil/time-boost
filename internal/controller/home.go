package controller

import (
	"time_boost/view/home"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Home(c echo.Context) error {
	return render(c, home.Home())
}
