package controllers

import (
	"log"
	"time"

	"github.com/aliffathurrisqi/GoFiber-MyApp/database"
	"github.com/aliffathurrisqi/GoFiber-MyApp/models"
	"github.com/aliffathurrisqi/GoFiber-MyApp/models/request"
	"github.com/gofiber/fiber/v2"
)

func UserAll(c *fiber.Ctx) error{

	var users []models.User

	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

    return c.JSON(users)
}

func UserCreate(c *fiber.Ctx) error{

	user := new(request.UserCreate)
	if err := c.BodyParser(user); err != nil{
		return err
	}

	createUser := models.User{
		Name : user.Name,
		Email: user.Email,
		City : user.City,
		Province: user.Province,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	request := database.DB.Debug().Create(&createUser)

	if request.Error != nil{
		return c.Status(500).JSON(fiber.Map{
			"message" : "failed"})
	}

	return c.JSON(fiber.Map{
		"message" : "success",
		"data" : createUser,
	})
}