package handler_test

import (
	"encoding/json"
	"errors"
	"myapp/internal/handler"
	"myapp/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) GetUsers() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func TestUserHandler_GetUsers(t *testing.T) {
	t.Run("ユースケース内部エラー", func(t *testing.T) {
		mockUsecase := new(MockUserUsecase)
		mockUsecase.On("GetUsers").Return([]model.User{}, errors.New("internal server error"))

		handler := handler.NewUserHandler(mockUsecase)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.GetUsers(c)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusInternalServerError, err.(*echo.HTTPError).Code)
		}
	})

	t.Run("成功", func(t *testing.T) {
		users := []model.User{
			{UserID: "user1", Email: "test1@test.com"},
			{UserID: "user2", Email: "test2@test.com"},
		}
		mockUsecase := new(MockUserUsecase)
		mockUsecase.On("GetUsers").Return(users, nil)

		handler := handler.NewUserHandler(mockUsecase)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.GetUsers(c)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
			expected := model.GetUsersResponse{
				Users: users,
			}
			expectedJSON, err := json.Marshal(expected)
			assert.NoError(t, err)
			assert.JSONEq(t, string(expectedJSON), rec.Body.String())
		}
	})
}
