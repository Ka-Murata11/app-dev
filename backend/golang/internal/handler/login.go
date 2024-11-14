package handler

import (
	"myapp/auth"
	"myapp/db"
	"myapp/entity"
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(e echo.Context) error {

	var loginRequest model.LoginRequest
	if err := e.Bind(&loginRequest); err != nil {
		return echo.ErrBadRequest
	}
	if err := e.Validate(&loginRequest); err != nil {
		return echo.ErrBadRequest
	}

	db, err := db.Init()
	if err != nil {
		return echo.ErrInternalServerError
	}

	var user entity.User
	if err := db.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		return echo.ErrUnauthorized
	}

	token, err := auth.CreateToken(user.ID, user.Name)
	if err != nil {
		return echo.ErrInternalServerError
	}

	auth.SetCookie(e, token)

	return e.JSON(http.StatusOK, "Login success")
}
