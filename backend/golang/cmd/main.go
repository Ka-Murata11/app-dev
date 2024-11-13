package main

import (
	"myapp/db"
	"myapp/entity"
	"myapp/internal/handler"
	"myapp/internal/repository"
	"myapp/internal/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := db.Init()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	e.GET("/users", userHandler.GetUsers)

	e.Logger.Fatal(e.Start(":1323"))
}
