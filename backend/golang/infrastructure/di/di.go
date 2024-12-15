package di

import (
	"myapp/db"
	"myapp/internal/handler"
	"myapp/internal/repository"
	"myapp/internal/usecase"
	"myapp/internal/util"
)

func InitializeLoginHandler() handler.LoginHandler {
	db, err := db.Init()
	if err != nil {
		panic(err)
	}

	util := util.NewPasswordUtil()
	r := repository.NewUserRepository(db)
	u := usecase.NewLoginUsecase(r, util)
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

func InitializeTaskHandler() handler.TaskHandler {
	db, err := db.Init()
	if err != nil {
		panic(err)
	}

	r := repository.NewTaskRepository(db)
	u := usecase.NewTaskUsecase(r)
	h := handler.NewTaskHandler(u)

	return h
}
