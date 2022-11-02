package main

import (
	"github.com/aliffathurrisqi/GoFiber-MyApp/database"
	migration "github.com/aliffathurrisqi/GoFiber-MyApp/database/migrations"
	Routes "github.com/aliffathurrisqi/GoFiber-MyApp/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	
	database.Connection()
	migration.Migration()

    app := fiber.New()

	Routes.Routes(app)

    app.Listen(":3000")
}