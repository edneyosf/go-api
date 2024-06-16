package controller

import (
	"net/http"
	"strconv"

	"base-api/internal/app/usecase"
	"github.com/labstack/echo/v4"
)

type BaseController struct {
	baseUsecase usecase.BaseUsecase
}

func NewController(baseUsecase usecase.BaseUsecase) *BaseController {
	return &BaseController{
		baseUsecase: baseUsecase,
	}
}

func (c *BaseController) Create(ctx echo.Context) error {
	var base struct {
		Name string `json:"name"`
	}
	base.Name = "sdf"
	if err := ctx.Bind(&base); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	newUser, err := c.baseUsecase.Create(base.Name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, newUser)
}

func (c *BaseController) GetById(ctx echo.Context) error {
	idParam := ctx.QueryParam("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user ID"})
	}

	user, err := c.baseUsecase.GetById(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, user)
}

func (c *BaseController) List(ctx echo.Context) error {
	users, err := c.baseUsecase.List()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, users)
}
