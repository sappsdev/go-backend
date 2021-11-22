package auth

import (
	"backend/config"
	"backend/token"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Roles(roles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		jwt := c.Get("x-token")
		if len(jwt) > 0 {
			valid, user := validToken(roles, jwt)
			if valid {
				c.Locals("User", user)
				return c.Next()
			}
		}
		return c.Status(401).JSON(fiber.Map{"auth": true})
	}

}

func validToken(r []string, tk string) (bool, token.UserAuth) {
	userAuth := new(token.UserAuth)
	token, _ := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		userAuth.ID = claims["id"].(string)
		userAuth.Email = claims["email"].(string)
		userAuth.Rol = claims["rol"].(string)
		if contains(r, userAuth.Rol) {
			return true, *userAuth
		}
	}

	return false, *userAuth

}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
