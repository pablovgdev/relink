package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pablovgdev/relink/internal/connection"
	"github.com/pablovgdev/relink/internal/handlers"
)

func init() {
	err := connection.Init()
	if err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()
	e.GET("/*", handlers.Redirect)
	e.POST("/redirects", handlers.PostRedirect)
	e.Logger.Fatal(e.Start(":8080"))
}
