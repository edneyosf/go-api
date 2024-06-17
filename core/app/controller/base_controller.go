package controller

import (
	"net/http"
	"strconv"

	"base-api/core/app/usecase"
	"github.com/labstack/echo/v4"
)

func (c *baseController) CreateBase(ctx echo.Context) error {
	var base struct {
		Name string `json:"name"`
	}
	if err := ctx.Bind(&base); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := c.baseUsecase.CreateBase(base.Name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, nil)
}

func (c *baseController) GetBaseById(ctx echo.Context) error {
	idParam := ctx.QueryParam("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user ID"})
	}

	user, err := c.baseUsecase.GetBaseById(id)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, user)
}

func (c *baseController) ListBases(ctx echo.Context) error {
	statusCode := http.StatusOK
	var data interface{}
	users, err := c.baseUsecase.ListBases()

	if err != nil {
		statusCode = http.StatusInternalServerError
		data = map[string]string{"error": err.Error()}
	} else{
		data = users
	}

	return ctx.JSON(statusCode, data)
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

