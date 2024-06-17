package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

const bases = "/bases"

func SetupRoutes(e *echo.Echo, db *gorm.DB) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	setupBase(e, db)
}
