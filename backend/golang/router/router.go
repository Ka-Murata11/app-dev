package router

import (
	"myapp/di"
	"myapp/internal/authMiddleware"
	"myapp/internal/handler"
	"myapp/validate"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {

	e.Validator = validate.NewValidator()

	// 認証が不要なAPI
	loginHandler := di.InitializeLoginHandler()
	e.POST("/signup", loginHandler.SignUp)
	e.POST("/signin", loginHandler.SignIn)
	e.GET("/user", handler.GetUser)

	// 認証が必要なAPI
	auth := e.Group("/api")
	auth.Use(authMiddleware.AuthMiddleware)

	userHandler := di.InitializeUserHandler()
	auth.GET("/users", userHandler.GetUsers)
}
