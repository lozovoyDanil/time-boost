package main

import (
	"time_boost/internal/controller"
)

func main() {
	handler := controller.NewHandler()
	app := handler.InitRoutes()

	err := app.Listen(":8080")
	if err != nil {
		app.Shutdown()
		panic(err)
	}
}
