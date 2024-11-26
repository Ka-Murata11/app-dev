package db

import "myapp/infrastructure/entity"

func Migrate() {
	db, err := Init()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.Task{})
	if err != nil {
		panic(err)
	}
}
