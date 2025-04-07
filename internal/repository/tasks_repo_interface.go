package repository

import "itmrchow/todolist-task/internal/entity"

type TasksRepository interface {
	CreateTask(task *entity.Task) error
	GetTask(id string) (*entity.Task, error)
	UpdateTask(task *entity.Task) error
	DeleteTask(id string) error
}
