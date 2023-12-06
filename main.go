package main

import (
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()
	app.GET("/start", StartHandler)
	app.Start()
}

func StartHandler(c *gofr.Context) (interface{}, error) {
	return "Welcome to the Project!", nil
}
