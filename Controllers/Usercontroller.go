package Controllers

import "github.com/gofiber/fiber/v2"

func All(c fiber.Ctx) error{
        return c.JSON(
			fiber.Map{
				"username" : "aliffathurrisqi", "password" : "password",
			})
    
	return nil
}