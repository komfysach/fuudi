package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/komfysach/fuudi/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
}
