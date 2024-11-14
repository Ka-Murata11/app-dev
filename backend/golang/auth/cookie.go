package auth

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func SetCookie(c echo.Context, token string) {
	expirationTime := time.Now().Add(30 * time.Minute)
	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Expires:  expirationTime,
		Path:     "/dashboard",
		HttpOnly: true,
	}

	c.SetCookie(cookie)
}
