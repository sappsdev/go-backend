package main

import (
	"fmt"
	"os"

	"backend/config"
	"backend/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	debug := config.Debug()
	if debug == false {
		config.GenSecret()
	}
	database.Connect()
	defer database.Cancel()
	defer database.Client.Disconnect(database.Ctx)

	database.Collections()

	_ = os.Mkdir(config.FilesAvatar(), os.ModePerm)

	app := fiber.New()

	app.Use(cors.New())
	app.Static("/files", config.FilesDir())

	Routes(app)
	app.Listen(fmt.Sprintf(":%s", config.Port()))
}