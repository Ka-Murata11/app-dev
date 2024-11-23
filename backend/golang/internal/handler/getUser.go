package handler

import (
	"myapp/db"
	"myapp/entity"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userRequest struct {
	UserID string `query:"user_id" validate:"required"`
}

type userResponse struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Job    string `json:"job"`
}

func GetUser(c echo.Context) error {
	var userRequest userRequest
	err := c.Bind(&userRequest)
	if err != nil {
		return echo.ErrBadRequest
	}

	if err := c.Validate(&userRequest); err != nil {
		return echo.ErrBadRequest
	}

	db, err := db.Init()
	if err != nil {
		return echo.ErrInternalServerError
	}

	var user entity.User
	if err := db.Where("user_id = ?", userRequest.UserID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.ErrBadRequest
		} else {
			return echo.ErrInternalServerError
		}
	}

	res := userResponse{
		UserID: user.UserID,
		Email:  user.Email,
		Role:   user.Role,
		Job:    user.Job,
	}
	return c.JSON(http.StatusOK, res)
}
