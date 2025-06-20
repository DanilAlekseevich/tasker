package controllers

import "tasker/internal/repositories"

type Controllers struct {
	Task TaskController
}

func NewControllers(repositories *repositories.Repositories) *Controllers {
	return &Controllers{
		Task: NewTaskController(&repositories.Task),
	}
}
