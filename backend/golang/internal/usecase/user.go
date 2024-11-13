package usecase

import (
	"myapp/entity"
	"myapp/internal/repository"
)

type UserUsecase interface {
	GetUsers() ([]entity.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) GetUsers() ([]entity.User, error) {
	users, err := u.userRepo.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
