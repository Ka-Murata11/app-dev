package usecase

import (
	"myapp/internal/model"
	"myapp/internal/repository"
)

type UserUsecase interface {
	GetUsers() ([]model.User, error)
	UpdateUser(updateUserInf model.UpdateUserInf) error
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

func (u *userUsecase) UpdateUser(updateUserInf model.UpdateUserInf) error {
	user, err := u.userRepo.GetUserByUserID(updateUserInf.UserID)
	if err != nil {
		return err
	}

	user.Email = updateUserInf.Email
	user.Job = updateUserInf.Job
	user.Role = updateUserInf.Role

	if err := u.userRepo.UpdateUser(user); err != nil {
		return err
	}

	return nil

}
