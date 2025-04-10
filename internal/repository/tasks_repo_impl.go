package repository

import (
	"context"

	"gorm.io/gorm"

	"itmrchow/todolist-task/internal/entity"
)

var _ TasksRepository = &database{}

type database struct {
	conn *gorm.DB
}

func NewTasksRepository(conn *gorm.DB) TasksRepository {
	return &database{
		conn: conn,
	}
}

func (d *database) CreateTask(ctx context.Context, task *entity.Task) error {
	return d.conn.WithContext(ctx).Create(task).Error
}

func (d *database) GetTask(ctx context.Context, id string, preloadOptions PreloadOption) (task *entity.Task, err error) {
	task = &entity.Task{}
	if preloadOptions.WithTaskList {
		err = d.conn.WithContext(ctx).Preload("TaskList").First(task, "id = ?", id).Error
	} else {
		err = d.conn.WithContext(ctx).First(task, "id = ?", id).Error
	}

	return task, err
}

func (d *database) UpdateTask(ctx context.Context, task *entity.Task) error {
	panic("TODO: Implement")
}

func (d *database) DeleteTask(ctx context.Context, id string) error {
	panic("TODO: Implement")
}

func (d *database) FindTask(ctx context.Context, page *entity.PageReqInfo) ([]*entity.Task, *entity.PageRespInfo, error) {
	panic("TODO: Implement")
}
