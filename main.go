package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, from Ostad! <3")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})
	e.GET("/k", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, kawsar from Yetfix! <3")
	})

	httpPort := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + httpPort))
}
