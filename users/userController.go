package users

import (
	"backend/token"

	"github.com/gofiber/fiber/v2"
)

func RequestPin(c *fiber.Ctx) error {

	userNew := new(UserNew)
	if err := c.BodyParser(userNew); err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "data": err})
	}

	if err := UniqueEmail(userNew.Email); err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "data": err})
	}

	return c.Status(400).JSON(fiber.Map{"success": true, "message": "send code"})
}

func Create(c *fiber.Ctx) error {

	userNew := new(UserNew)

	if err := c.BodyParser(userNew); err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "data": err})
	}

	rol := "user"

	userId, err := Insert(*userNew, rol)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "message": err})
	}

	userAuth := new(token.UserAuth)
	userAuth.ID = userId
	userAuth.Email = userNew.Email
	userAuth.Rol = rol

	token := token.JwtNew(*userAuth)

	return c.Status(200).JSON(fiber.Map{"success": true, "message": "user created", "data": token})
}
