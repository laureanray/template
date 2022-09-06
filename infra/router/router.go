package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"template/interface/controller"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(context echo.Context) error { return c.User.GetUsers(context) })
	e.POST("/users", func(context echo.Context) error { return c.User.CreateUser(context) })

	return e
}
