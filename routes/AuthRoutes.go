package routes

import (
	"deka_backend/controllers"

	"github.com/gofiber/fiber"
)

func UsersAPI(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/users", controllers.GetAllUsersController)
	api.Get("/users/:id", controllers.GetUserByIDController)
	api.Post("/users", controllers.CreateUserController)
	api.Delete("/users/:id", controllers.DeleteUserController)
}
