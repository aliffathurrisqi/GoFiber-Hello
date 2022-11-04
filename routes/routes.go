package routes

import (
	"github.com/aliffathurrisqi/GoFiber-MyApp/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App){

	app.Get("/user", controllers.UserAll)

	app.Post("/user", controllers.UserCreate)


}