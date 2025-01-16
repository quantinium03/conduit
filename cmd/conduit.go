package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/quantinium03/conduit/internal/config"
	"github.com/quantinium03/conduit/internal/metadata"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	metadata, err := metadata.StartService()
	if err != nil {
		e.Logger.Fatal(err)
		return
	}

	g := e.Group(config.Settings.RoutePrefix)
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
