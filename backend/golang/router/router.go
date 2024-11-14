package router

import (
	"myapp/authMiddleware"
	"myapp/di"
	"myapp/internal/handler"
	"myapp/validate"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {

	e.Validator = validate.NewValidator()

	e.GET("/login", handler.Login)

	auth := e.Group("/api")

	auth.Use(authMiddleware.AuthMiddleware)

	h := di.InitializeUserHandler()
	auth.GET("/users", h.GetUsers)
}
