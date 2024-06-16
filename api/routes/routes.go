package routes

import (
	"base-api/internal/app/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(e *echo.Echo, userController *controller.BaseController) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", userController.Create)
	e.GET("/users", userController.List)
	e.GET("/users/:id", userController.GetById)
}
