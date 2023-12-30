package main

import (
	"time_boost/internal/controller"
)

func main() {
	handler := controller.NewHandler()
	app := handler.InitRoutes()

	app.Listen(":8080")

}
