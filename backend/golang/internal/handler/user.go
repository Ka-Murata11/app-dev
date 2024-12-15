package handler

import (
	"myapp/internal/model"
	"myapp/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandler interface {
	GetUsers(c echo.Context) error
	UpdateUser(c echo.Context) error
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

func (h *userHandler) UpdateUser(c echo.Context) error {
	var updateUserInf model.UpdateUserInf
	if err := c.Bind(&updateUserInf); err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(&updateUserInf); err != nil {
		return echo.ErrBadRequest
	}

	if err := h.userUsecase.UpdateUser(updateUserInf); err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.ErrBadRequest
		} else {
			return echo.ErrInternalServerError
		}
	}

	return c.JSON(http.StatusOK, "success user update")
}
