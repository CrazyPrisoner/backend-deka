package routes

import (
	"github.com/gofiber/fiber"
)

func Routing(app *fiber.App) {
	UsersAPI(app)
}
