package routes

import (
	"tasker/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, controllers *controllers.Controllers) {
	SetupTaskRouters(app, controllers.Task)
}
