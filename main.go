package main

import (
	"github.com/aliffathurrisqi/GoFiber-MyApp/Routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

	Routes.RouteInit(app)

    app.Listen(":3000")
}