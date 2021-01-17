package main

import (
	"github.com/RisingGeek/messy-stuff/delivery"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", delivery.ChannelInfo)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

}
