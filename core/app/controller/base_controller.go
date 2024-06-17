package controller

import (
	"base-api/core/app/data/request"
	"base-api/core/app/message"
	"net/http"
	"strconv"

	"base-api/core/app/usecase"
	"github.com/labstack/echo/v4"
)

func (controller *baseController) CreateBase(ctx echo.Context) error {
	var base request.BaseData

	err := ctx.Bind(&base)

	if err != nil {
		response := errorResponse(message.FailureToBindData)

		return ctx.JSON(http.StatusBadRequest, response)
	}

	err = controller.baseUsecase.CreateBase(base)

	if err != nil {
		response := errorResponse(err.Error())

		return ctx.JSON(http.StatusInternalServerError, response)
	}

	return ctx.NoContent(http.StatusCreated)
}

func (controller *baseController) GetBaseById(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		response := errorResponse(message.FailureToBindData)

		return ctx.JSON(http.StatusBadRequest, response)
	}

	base, err := controller.baseUsecase.GetBaseById(id)

	if err != nil {
		response := errorResponse(err.Error())

		return ctx.JSON(http.StatusNotFound, response)
	}

	return ctx.JSON(http.StatusOK, base)
}

func (controller *baseController) ListBases(ctx echo.Context) error {
	var response interface{}
	statusCode := http.StatusOK
	bases, err := controller.baseUsecase.ListBases()

	if err != nil {
		statusCode = http.StatusInternalServerError
		response = errorResponse(err.Error())
	} else{
		response = bases
	}

	return ctx.JSON(statusCode, response)
}

func NewController(baseUsecase usecase.BaseUsecase) BaseController {
	return &baseController{ baseUsecase }
}

type baseController struct { baseUsecase usecase.BaseUsecase }

type BaseController interface {
	CreateBase(ctx echo.Context) error
	GetBaseById(ctx echo.Context) error
	ListBases(ctx echo.Context) error
}

