package main

import (
	"github.com/preritshrivastava2002/Management-System/handler"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()
	app.GET("/start", handler.StartHandler)
	app.GET("/homes", handler.AllHomeHandler)
	app.POST("/homes", handler.CreateHomeHandler)
	app.GET("/homes/{id}", handler.GetHomeHandler)
	app.PUT("/homes/{id}", handler.UpdateHomeHandler)
	app.DELETE("/homes/{id}", handler.DeleteHomeHandler)
	app.Start()
}
