package main

import (
	"app/src/interfaces"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET},
	}))

	interfaces.Routes(e)
	interfaces.Run(e, os.Getenv("BE_PORT"))
}
