package controllers

import (
	"log"
	"time"

	"github.com/aliffathurrisqi/GoFiber-MyApp/database"
	"github.com/aliffathurrisqi/GoFiber-MyApp/models"
	"github.com/aliffathurrisqi/GoFiber-MyApp/models/request"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func UserAll(c *fiber.Ctx) error{

	var users []models.User

	result := database.DB.Debug().Order("name").Find(&users)

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

	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil{
		return c.Status(400).JSON(fiber.Map{
			"message" : "failed",
			"error" : errValidate.Error(),
		})
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

func UserFind(c *fiber.Ctx) error{

	userId := c.Params("id")

	var user models.User

	q := database.DB.First(&user, "id = ?" , userId)

	if q.Error != nil{
		return c.Status(400).JSON(fiber.Map{
			"message" : "user not found",
		})
	}

	return c.JSON(fiber.Map{
		"message" : "success",
		"data" : fiber.Map{
			"name" : user.Name, 
			"email" : user.Email,
			"city" : user.City,
			"province" : user.Province,
			},
	})
}

func UserEdit(c *fiber.Ctx) error{

	userRequest := new(request.UserEdit)

	if err := c.BodyParser(userRequest); err != nil{
		return c.Status(400).JSON(fiber.Map{
			"message" : "bad request",
		})
	}

	var user models.User


	userId := c.Params("id")

	q := database.DB.First(&user, "id = ?" , userId)

	if q.Error != nil{
		return c.Status(400).JSON(fiber.Map{
			"message" : "user not found",
		})
	}

	if userRequest.Name != ""{
		user.Name = userRequest.Name
	}

	if userRequest.City != ""{
		user.City = userRequest.City
	}

	if userRequest.Province != ""{
		user.Province = userRequest.Province
	}

	user.Updated_at = time.Now()

	update := database.DB.Save(&user)

	if update.Error != nil{
		return c.Status(500).JSON(fiber.Map{
			"message" : "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message" : "success",
		"data" : user,
	})
	
}

func UserDelete(c *fiber.Ctx) error{

	userId := c.Params("id")

	var user models.User

	q := database.DB.First(&user, "id = ?" , userId)

	if q.Error != nil{
		return c.Status(404).JSON(fiber.Map{
			"message" : "user not found",
		})
	}

	del := database.DB.Delete(&user)

	if del.Error != nil{
		return c.Status(500).JSON(fiber.Map{
			"message" : "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message" : "user was deleted",
	})

}