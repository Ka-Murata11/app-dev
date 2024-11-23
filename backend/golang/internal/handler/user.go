package handler

import (
	"myapp/internal/usecase"
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetUsers(c echo.Context) error
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase}
}

func (h *userHandler) GetUsers(c echo.Context) error {
	users, err := h.userUsecase.GetUsers()
	if err != nil {
		return echo.ErrInternalServerError
	}

	res := model.GetUsersResponse{
		Users: users,
	}

	return c.JSON(http.StatusOK, res)
}
