package repository

import (
	"context"
	"fmt"
	"math"

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

func (d *database) UpdateTask(ctx context.Context, id string, task map[string]any) error {
	return d.conn.WithContext(ctx).Model(&entity.Task{}).Where("id = ?", id).Updates(task).Error
}

func (d *database) DeleteTask(ctx context.Context, id string) error {
	return d.conn.WithContext(ctx).Delete(&entity.Task{}, "id = ?", id).Error
}

type FindTaskParams struct {
	UserID string
	TaskID *string
	Status *entity.TaskStatus
}

func (d *database) FindTask(ctx context.Context, params FindTaskParams, pageReq *entity.PageReqInfo) (tasks []*entity.Task, pageInfo *entity.PageRespInfo, err error) {
	// filter
	query := d.conn.Model(&entity.Task{}).Where("user_id = ?", params.UserID)

	if params.TaskID != nil {
		query = query.Where("task_list_id = ?", params.TaskID)
	} else {
		query = query.Where("task_list_id IS NULL")
	}

	if params.Status != nil {
		query = query.Where("status = ?", params.Status)
	}

	// order
	for _, sort := range pageReq.Sort {
		query = query.Order(fmt.Sprintf("%s %s", sort.Property, sort.Direction))
	}

	// pagination
	query, pageInfo, err = paginate(entity.Task{}, pageReq, query)
	if err != nil {
		return nil, nil, err
	}

	// execute query
	if err := query.Find(&tasks).Error; err != nil {
		return nil, nil, err
	}

	return
}

func paginate[T any](model T, pageReq *entity.PageReqInfo, db *gorm.DB) (*gorm.DB, *entity.PageRespInfo, error) {
	// init page resp
	pageResp := &entity.PageRespInfo{
		Page:  pageReq.Page,
		Limit: pageReq.Limit,
	}

	// total
	var total int64
	if err := db.Model(model).Count(&total).Error; err != nil {
		return nil, nil, err
	}
	pageResp.Total = total

	// total pages
	pageResp.TotalPages = int64(math.Ceil(float64(total) / float64(pageResp.Limit)))

	// offset
	offset := (pageReq.Page - 1) * pageReq.Limit

	// query
	query := db.Offset(int(offset)).Limit(int(pageReq.Limit))

	return query, pageResp, nil
}
