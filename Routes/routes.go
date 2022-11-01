package Routes

import "github.com/gofiber/fiber/v2"

func RouteInit(app *fiber.App){

	app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(
			fiber.Map{
				"name" : "Aliffathur Risqi", "age" : 22, "job" : "Web Programmer",
			})
    })

}