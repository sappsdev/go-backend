package token

import (
	"time"

	"backend/config"

	"github.com/golang-jwt/jwt"
)

func JwtNew(user UserAuth) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"rol":   user.Rol,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(config.SECRET))
	return tokenString
}