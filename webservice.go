package main

import (
	"coinConversion/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func webServer(startPort string) {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	r := e.Group("/exchange")
	r.GET("/:amount/:from/:to/:rate", controller.GetExchange)

	c := e.Group("/consults")
	c.GET("", controller.GetConsults)

	e.Logger.Fatal(e.Start(startPort))
}
