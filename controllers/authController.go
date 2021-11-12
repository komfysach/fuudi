package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/komfysach/fuudi/models"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "Sachin",
	}

	user.LastName = "Lendis"

	return c.JSON(user)
}
