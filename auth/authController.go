package auth

import (
	"backend/token"
	"backend/users"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {

	login := new(users.UserLogin)

	authResponse := new(AuthResponse)

	if err := c.BodyParser(login); err != nil {
		authResponse.Errors.Email = "Error en envío de los datos"
		return c.Status(400).JSON(authResponse)
	}

	user, err := users.FindByEmail(login.Email)
	if err != nil {
		authResponse.Errors.Email = "Usuario o contraseña incorrecta"
		return c.Status(400).JSON(authResponse)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		authResponse.Errors.Password = "Error en envío de los datos"
		return c.Status(400).JSON(authResponse)
	}

	userId := user.ID.Hex()

	userAuth := new(token.UserAuth)
	userAuth.ID = userId
	userAuth.Email = user.Email
	userAuth.Rol = user.Rol

	token := token.JwtNew(*userAuth)
	authResponse.Rol = user.Rol
	authResponse.Token = token

	return c.Status(200).JSON(authResponse)

}

func CheckAuth(c *fiber.Ctx) error {
	authUser := c.Locals("User").(token.UserAuth)
	authResponse := new(AuthResponse)
	authResponse.Rol = authUser.Rol
	return c.Status(200).JSON(authResponse)
	
}
