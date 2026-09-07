package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/controllers"
)

func UserRoutes(
	app *fiber.App,
	controller *controllers.UserController,
) {

	users := app.Group("/api/users")

	users.Get("/", controller.GetUsers)

	users.Get("/:id", controller.GetUserByID)

	users.Post("/", controller.CreateUser)

	users.Put("/:id", controller.UpdateUser)

	users.Delete("/:id", controller.DeleteUser)
}
