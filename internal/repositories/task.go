package repositories

import (
	"context"
	"database/sql"
	"tasker/internal/model"
)

type TaskRepository interface {
	Create(ctx context.Context, task *model.Task) error
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(ctx context.Context, task *model.Task) error {
	query := `
        INSERT INTO tasks (id, name, email)
        VALUES ($1, $2, $3)
        RETURNING id
    `

	return r.db.QueryRowContext(ctx, query, task.Id, task.Name, task.Description).Scan(&task.Id)
}
