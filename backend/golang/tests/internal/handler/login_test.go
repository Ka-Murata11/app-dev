package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"myapp/infrastructure/entity"
	"myapp/internal/handler"
	"myapp/internal/model"
	"myapp/validate"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLoginUsecase struct {
	mock.Mock
}

func (m *MockLoginUsecase) SignIn(loginRequest model.SignInRequest) (entity.User, error) {
	args := m.Called(loginRequest)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockLoginUsecase) SignUp(signUpRequest model.SignUpRequest) error {
	args := m.Called(signUpRequest)
	return args.Error(0)
}

func TestLoginHandler_SignIn(t *testing.T) {
	t.Run("リクエスト不正 リクエストボディに値が無い", func(t *testing.T) {
		handler := handler.NewLoginHandler(nil)

		e := echo.New()
		e.Validator = validate.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.SignIn(c)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
		}
	})

	t.Run("リクエスト不正 ユーザIDが無い", func(t *testing.T) {
		handler := handler.NewLoginHandler(nil)

		loginRequest := model.SignInRequest{UserID: "", Password: "password"}
		reqBody, err := json.Marshal(loginRequest)
		assert.NoError(t, err)

		e := echo.New()
		e.Validator = validate.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = handler.SignIn(c)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
		}
	})

	t.Run("リクエスト不正 パスワードが無い", func(t *testing.T) {
		handler := handler.NewLoginHandler(nil)

		loginRequest := model.SignInRequest{UserID: "test", Password: ""}
		reqBody, err := json.Marshal(loginRequest)
		assert.NoError(t, err)

		e := echo.New()
		e.Validator = validate.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = handler.SignIn(c)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
		}
	})

	t.Run("ユースケース内部エラー", func(t *testing.T) {
		loginRequest := model.SignInRequest{UserID: "test", Password: "password"}
		reqBody, err := json.Marshal(loginRequest)
		assert.NoError(t, err)

		MockUsecase := new(MockLoginUsecase)
		MockUsecase.On("SignIn", loginRequest).Return(entity.User{}, errors.New("internal server error"))

		handler := handler.NewLoginHandler(MockUsecase)

		e := echo.New()
		e.Validator = validate.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = handler.SignIn(c)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusInternalServerError, err.(*echo.HTTPError).Code)
		}
	})

	t.Run("認証エラー", func(t *testing.T) {
		loginRequest := model.SignInRequest{UserID: "test", Password: "password"}
		reqBody, err := json.Marshal(loginRequest)
		assert.NoError(t, err)

		MockUsecase := new(MockLoginUsecase)
		MockUsecase.On("SignIn", loginRequest).Return(entity.User{}, errors.New("user not found"))

		handler := handler.NewLoginHandler(MockUsecase)

		e := echo.New()
		e.Validator = validate.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = handler.SignIn(c)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusUnauthorized, err.(*echo.HTTPError).Code)
		}
	})

	t.Run("成功", func(t *testing.T) {
		loginRequest := model.SignInRequest{UserID: "test", Password: "password"}
		reqBody, err := json.Marshal(loginRequest)
		assert.NoError(t, err)

		user := entity.User{
			UserID:   "test",
			Password: "password",
		}

		MockUsecase := new(MockLoginUsecase)
		MockUsecase.On("SignIn", loginRequest).Return(user, nil)

		handler := handler.NewLoginHandler(MockUsecase)

		e := echo.New()
		e.Validator = validate.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = handler.SignIn(c)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
			cookie := rec.Result().Cookies()
			assert.NotEmpty(t, cookie)
			assert.Equal(t, "Authorization", cookie[0].Name)
			assert.NotEmpty(t, cookie[0].Value)
		}
	})
}

func TestLoginHandler_SignUp(t *testing.T) {
	t.Run("リクエスト不正 リクエストボディに値が無い", func(t *testing.T) {
		handler := handler.NewLoginHandler(nil)

		e := echo.New()
		e.Validator = validate.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.SignUp(c)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
		}
	})

	t.Run("リクエスト不正 ユーザIDが無い", func(t *testing.T) {
		handler := handler.NewLoginHandler(nil)

		loginRequest := model.SignUpRequest{UserID: "", Password: "password", Email: "test@test.com"}
		reqBody, err := json.Marshal(loginRequest)
		assert.NoError(t, err)

		e := echo.New()
		e.Validator = validate.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = handler.SignUp(c)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
		}
	})

	t.Run("リクエスト不正 パスワードが無い", func(t *testing.T) {
		handler := handler.NewLoginHandler(nil)

		loginRequest := model.SignUpRequest{UserID: "test", Password: "", Email: "test@test.com"}
		reqBody, err := json.Marshal(loginRequest)
		assert.NoError(t, err)

		e := echo.New()
		e.Validator = validate.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = handler.SignUp(c)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
		}
	})

	t.Run("リクエスト不正 メールアドレスが無い", func(t *testing.T) {
		handler := handler.NewLoginHandler(nil)

		loginRequest := model.SignUpRequest{UserID: "test", Password: "password", Email: ""}
		reqBody, err := json.Marshal(loginRequest)
		assert.NoError(t, err)

		e := echo.New()
		e.Validator = validate.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = handler.SignUp(c)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
		}
	})

	t.Run("リクエスト不正 メールアドレスの形式が違う", func(t *testing.T) {
		handler := handler.NewLoginHandler(nil)

		loginRequest := model.SignUpRequest{UserID: "test", Password: "password", Email: "test"}
		reqBody, err := json.Marshal(loginRequest)
		assert.NoError(t, err)

		e := echo.New()
		e.Validator = validate.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = handler.SignUp(c)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
		}
	})

	t.Run("ユーザIDもしくはメールアドレスが重複している", func(t *testing.T) {
		signUpRequest := model.SignUpRequest{UserID: "test", Password: "password", Email: "test@test.com"}
		reqBody, err := json.Marshal(signUpRequest)
		assert.NoError(t, err)

		MockUsecase := new(MockLoginUsecase)
		MockUsecase.On("SignUp", signUpRequest).Return(errors.New("user already exists"))

		handler := handler.NewLoginHandler(MockUsecase)

		e := echo.New()
		e.Validator = validate.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = handler.SignUp(c)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
		}
	})

	t.Run("ユーザIDもしくはメールアドレスが重複している", func(t *testing.T) {
		signUpRequest := model.SignUpRequest{UserID: "test", Password: "password", Email: "test@test.com"}
		reqBody, err := json.Marshal(signUpRequest)
		assert.NoError(t, err)

		MockUsecase := new(MockLoginUsecase)
		MockUsecase.On("SignUp", signUpRequest).Return(errors.New("user already exists"))

		handler := handler.NewLoginHandler(MockUsecase)

		e := echo.New()
		e.Validator = validate.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = handler.SignUp(c)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
		}
	})

	t.Run("ユースケース 内部エラー", func(t *testing.T) {
		signUpRequest := model.SignUpRequest{UserID: "test", Password: "password", Email: "test@test.com"}
		reqBody, err := json.Marshal(signUpRequest)
		assert.NoError(t, err)

		MockUsecase := new(MockLoginUsecase)
		MockUsecase.On("SignUp", signUpRequest).Return(errors.New("internal server error"))

		handler := handler.NewLoginHandler(MockUsecase)

		e := echo.New()
		e.Validator = validate.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = handler.SignUp(c)
		if assert.Error(t, err) {
			assert.Equal(t, http.StatusInternalServerError, err.(*echo.HTTPError).Code)
		}
	})

	t.Run("成功", func(t *testing.T) {
		signUpRequest := model.SignUpRequest{UserID: "test", Password: "password", Email: "test@test.com"}
		reqBody, err := json.Marshal(signUpRequest)
		assert.NoError(t, err)

		MockUsecase := new(MockLoginUsecase)
		MockUsecase.On("SignUp", signUpRequest).Return(nil)

		handler := handler.NewLoginHandler(MockUsecase)

		e := echo.New()
		e.Validator = validate.NewValidator()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = handler.SignUp(c)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
