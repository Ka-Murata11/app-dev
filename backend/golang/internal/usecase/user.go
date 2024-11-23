package usecase

import (
	"myapp/internal/repository"
	"myapp/model"
)

type UserUsecase interface {
	GetUsers() ([]model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) GetUsers() ([]model.User, error) {
	users, err := u.userRepo.GetUsers()
	if err != nil {
		return nil, err
	}

	resUsers := make([]model.User, len(users))
	for i, user := range users {
		resUsers[i].UserID = user.UserID
		resUsers[i].Email = user.Email
		resUsers[i].Job = user.Job
		resUsers[i].Role = user.Role
	}

	return resUsers, nil
}
