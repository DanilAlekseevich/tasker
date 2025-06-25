package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log/slog"
	"tasker/internal/config"
	"tasker/internal/controllers"
	"tasker/internal/routes"
)

type ServerFactory struct {
	config *config.Config
	logger *slog.Logger
}

func New(cfg *config.Config, logger *slog.Logger) *ServerFactory {
	return &ServerFactory{
		config: cfg,
		logger: logger,
	}
}

func (sf *ServerFactory) CreateServer(controllers *controllers.Controllers) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: sf.errorHandler,
	})

	app.Use(logger.New())
	app.Use(cors.New())

	routes.Setup(app, controllers)

	return app
}

func (sf *ServerFactory) errorHandler(c *fiber.Ctx, err error) error {
	sf.logger.Error("request error", "error", err, "path", c.Path())

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Internal Server Error",
	})
}
