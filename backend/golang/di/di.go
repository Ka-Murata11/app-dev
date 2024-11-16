package di

import (
	"myapp/db"
	"myapp/internal/handler"
	"myapp/internal/repository"
	"myapp/internal/usecase"
)

func InitializeLoginHandler() handler.LoginHandler {
	db, err := db.Init()
	if err != nil {
		panic(err)
	}

	r := repository.NewUserRepository(db)
	u := usecase.NewLoginUsecase(r)
	h := handler.NewLoginHandler(u)

	return h
}

func InitializeUserHandler() handler.UserHandler {
	db, err := db.Init()
	if err != nil {
		panic(err)
	}

	r := repository.NewUserRepository(db)
	u := usecase.NewUserUsecase(r)
	h := handler.NewUserHandler(u)

	return h
}
