package repository

import (
	"context"

	"itmrchow/todolist-task/internal/entity"
)

type TasksRepository interface {
	CreateTask(ctx context.Context, task *entity.Task) error
	GetTask(ctx context.Context, id string, preloadOptions PreloadOption) (*entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id string) error
	FindTask(ctx context.Context, page *entity.PageReqInfo) ([]*entity.Task, *entity.PageRespInfo, error)
}

type PreloadOption struct {
	WithTaskList bool
}
