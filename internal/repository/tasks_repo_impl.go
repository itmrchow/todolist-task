package repository

import (
	"gorm.io/gorm"

	"itmrchow/todolist-task/internal/entity"
)

var _ TasksRepository = &database{}

type database struct {
	conn *gorm.DB
}

func (d *database) CreateTask(task *entity.Task) error {
	panic("TODO: Implement")
}

func (d *database) GetTask(id string) (*entity.Task, error) {
	panic("TODO: Implement")
}

func (d *database) UpdateTask(task *entity.Task) error {
	panic("TODO: Implement")
}

func (d *database) DeleteTask(id string) error {
	panic("TODO: Implement")
}

func NewTasksRepository(conn *gorm.DB) TasksRepository {
	return &database{
		conn: conn,
	}
}
