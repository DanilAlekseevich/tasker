package routes

import (
	"tasker/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupTaskRouters(app *fiber.App, controller controllers.TaskController) {
	task := app.Group("/task")
	task.Post("/create", controller.Create)
}
