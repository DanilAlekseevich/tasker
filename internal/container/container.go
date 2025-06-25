package container

import (
	"database/sql"
	"log/slog"
	"tasker/internal/config"
	"tasker/internal/controllers"
	"tasker/internal/repositories"
)

type Container struct {
	Config       *config.Config
	Logger       *slog.Logger
	DB           *sql.DB
	Repositories *repositories.Repositories
	Controllers  *controllers.Controllers
}

func New(cfg *config.Config, logger *slog.Logger, db *sql.DB) *Container {
	return &Container{
		Config: cfg,
		Logger: logger,
		DB:     db,
	}
}

func (c *Container) InitRepositories() error {
	c.Repositories = repositories.NewRepositories(c.DB)
	return nil
}

func (c *Container) InitControllers() error {
	c.Controllers = controllers.NewControllers(c.Repositories)
	return nil
}

func (c *Container) InitAll() error {
	if err := c.InitRepositories(); err != nil {
		return err
	}
	if err := c.InitControllers(); err != nil {
		return err
	}
	return nil
}
