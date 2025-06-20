package controllers

import (
	"tasker/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

type TaskController interface {
	Create(c *fiber.Ctx) error
}

type taskController struct {
	repo *repositories.TaskRepository
}

func NewTaskController(repo *repositories.TaskRepository) TaskController {
	return &taskController{repo: repo}
}

func (tc *taskController) Create(c *fiber.Ctx) error {
	return fiber.NewError(404, "Custom message")
}
