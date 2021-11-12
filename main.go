package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/komfysach/fuudi/database"
	"github.com/komfysach/fuudi/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
