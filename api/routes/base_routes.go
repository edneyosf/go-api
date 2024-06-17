package routes

import (
	"base-api/core/app/controller"
	"base-api/core/app/repository"
	"base-api/core/app/usecase"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func setupBase(e *echo.Echo, db *gorm.DB) {
	baseRepo := repository.NewBaseRepository(db)
	baseUsecase := usecase.NewBaseUsecase(baseRepo)
	baseController := controller.NewController(baseUsecase)

	e.POST(bases, baseController.CreateBase)
	e.GET(bases, baseController.ListBases)
	e.GET(bases+ "/:id", baseController.GetBaseById)
}


