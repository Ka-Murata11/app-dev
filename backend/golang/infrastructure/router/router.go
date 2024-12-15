package router

import (
	"myapp/infrastructure/di"
	"myapp/internal/authMiddleware"
	"myapp/internal/handler"
	"myapp/validate"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router(e *echo.Echo) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://frontend:3000", "http://localhost:3000", "http://frontend:1323", "http://localhost:1323"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Authorization"},
	}))

	e.Validator = validate.NewValidator()

	// 認証が不要なAPI
	loginHandler := di.InitializeLoginHandler()
	e.POST("/signup", loginHandler.SignUp)
	e.POST("/signin", loginHandler.SignIn)

	e.POST("/csv", handler.ImportCSV)

	// 認証が必要なAPI
	auth := e.Group("/api")
	auth.Use(authMiddleware.AuthMiddleware)

	userHandler := di.InitializeUserHandler()
	auth.GET("/users", userHandler.GetUsers)
}
