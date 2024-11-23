package usecase_test

import (
	"errors"
	"myapp/entity"
	"myapp/internal/usecase"
	"myapp/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUsers() ([]entity.User, error) {
	args := m.Called()
	return args.Get(0).([]entity.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByUserID(userID string) (entity.User, error) {
	args := m.Called(userID)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByEmail(email string) (entity.User, error) {
	args := m.Called(email)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) CreateUser(user entity.User) error {
	args := m.Called(user)
	return args.Error(1)
}

func TestUserUsecase_GetUsers(t *testing.T) {

	t.Run("リポジトリ内部エラー", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		mockRepo.On("GetUsers").Return([]entity.User{}, errors.New("internal server error"))

		usecase := usecase.NewUserUsecase(mockRepo)

		_, err := usecase.GetUsers()
		assert.Error(t, err)
	})

	t.Run("成功 0件取得", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		mockRepo.On("GetUsers").Return([]entity.User{}, nil)

		usecase := usecase.NewUserUsecase(mockRepo)

		actual, err := usecase.GetUsers()
		if assert.NoError(t, err) {
			expected := []model.User{}
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("成功 1件取得", func(t *testing.T) {
		users := []entity.User{
			{UserID: "user1", Email: "test1@test.com", Password: "password"},
		}
		mockRepo := new(MockUserRepository)
		mockRepo.On("GetUsers").Return(users, nil)

		usecase := usecase.NewUserUsecase(mockRepo)

		actual, err := usecase.GetUsers()
		if assert.NoError(t, err) {
			expected := []model.User{
				{UserID: "user1", Email: "test1@test.com"},
			}
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("成功 2件取得", func(t *testing.T) {
		users := []entity.User{
			{UserID: "user1", Email: "test1@test.com", Password: "password"},
			{UserID: "user2", Email: "test2@test.com", Password: "password"},
		}
		mockRepo := new(MockUserRepository)
		mockRepo.On("GetUsers").Return(users, nil)

		usecase := usecase.NewUserUsecase(mockRepo)

		actual, err := usecase.GetUsers()
		if assert.NoError(t, err) {
			expected := []model.User{
				{UserID: "user1", Email: "test1@test.com"},
				{UserID: "user2", Email: "test2@test.com"},
			}
			assert.Equal(t, expected, actual)
		}
	})
}
