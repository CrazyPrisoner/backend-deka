package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"

	"deka_backend/routes"
)

func Test(c *fiber.Ctx) {
	c.Send("Test Test Dead")
}

func main() {
	app := fiber.New()
	app.Use(middleware.Logger())
	app.Get("/", Test)
	routes.Routing(app)
	app.Listen(3000)
}
