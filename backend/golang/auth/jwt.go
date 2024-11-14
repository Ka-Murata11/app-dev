package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

type ResponceJWTTokenJSON struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

var jwtSecret = []byte("secret")

func CreateToken(id uint, name string) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	// expirationTime := time.Now()
	claims := &JwtCustomClaims{
		ID:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*JwtCustomClaims, error) {
	claims := &JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
