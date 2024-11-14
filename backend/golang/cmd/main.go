package main

import (
	"myapp/db"
	"myapp/entity"
	"myapp/router"

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

	router.Router(e)

	e.Logger.Fatal(e.Start(":1323"))
}
