package routes

import (
	"github.com/aliffathurrisqi/GoFiber-MyApp/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App){

	app.Get("/user", controllers.UserAll)

	app.Get("/user/:id", controllers.UserFind)

	app.Post("/user", controllers.UserCreate)


}