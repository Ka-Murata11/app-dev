package router

import (
	"myapp/infrastructure/di"
	"myapp/internal/authMiddleware"
	"myapp/validate"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router(e *echo.Echo) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = validate.NewValidator()

	// 認証が不要なAPI
	loginHandler := di.InitializeLoginHandler()
	e.POST("/signup", loginHandler.SignUp)
	e.POST("/signin", loginHandler.SignIn)

	// 認証が必要なAPI
	auth := e.Group("/api")
	auth.Use(authMiddleware.AuthMiddleware)

	userHandler := di.InitializeUserHandler()
	auth.GET("/users", userHandler.GetUsers)
}
