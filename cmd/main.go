package main

import (
	"context"
	"time_boost/internal/controller"
)

func main() {
	handler := controller.NewHandler()
	app := handler.InitRoutes()

	err := app.Start(":8080")
	if err != nil {
		app.Shutdown(context.Background())
		panic(err)
	}
}
