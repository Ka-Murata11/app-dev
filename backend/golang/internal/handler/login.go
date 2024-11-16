package handler

import (
	"myapp/auth"
	"myapp/internal/usecase"
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginHandler interface {
	SignIn(e echo.Context) error
	SignUp(e echo.Context) error
}

type loginhandler struct {
	loginUsecase usecase.LoginUsecase
}

func NewLoginHandler(loginUsecase usecase.LoginUsecase) LoginHandler {
	return &loginhandler{loginUsecase}
}

func (h *loginhandler) SignIn(e echo.Context) error {

	var loginRequest model.SignInRequest
	if err := e.Bind(&loginRequest); err != nil {
		return echo.ErrBadRequest
	}
	if err := e.Validate(&loginRequest); err != nil {
		return echo.ErrBadRequest
	}

	user, err := h.loginUsecase.SignIn(loginRequest)
	if err != nil {
		if err.Error() == "internal server error" {
			return echo.ErrInternalServerError
		} else {
			return echo.ErrUnauthorized
		}
	}

	token, err := auth.CreateToken(user.UserID, user.Role)
	if err != nil {
		return echo.ErrInternalServerError
	}

	auth.SetCookie(e, token)

	return e.JSON(http.StatusOK, "Login success")
}

func (h *loginhandler) SignUp(e echo.Context) error {
	var signUpRequest model.SignUpRequest
	if err := e.Bind(&signUpRequest); err != nil {
		return echo.ErrBadRequest
	}
	if err := e.Validate(&signUpRequest); err != nil {
		return echo.ErrBadRequest
	}

	err := h.loginUsecase.SignUp(signUpRequest)
	if err != nil {
		if err.Error() == "user already exists" {
			return echo.ErrBadRequest
		} else {
			return echo.ErrInternalServerError
		}
	}

	return e.JSON(http.StatusOK, "Sign up success")
}
