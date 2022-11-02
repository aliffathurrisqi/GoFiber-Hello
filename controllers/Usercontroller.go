package controllers

import "github.com/gofiber/fiber/v2"

func UserAll(c *fiber.Ctx) error{
        return c.JSON(fiber.Map{
				"username" : "aliffathurrisqi", "password" : "password",
			})
}