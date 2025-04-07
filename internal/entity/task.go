package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string

const (
	StatusPending    Status = "Pending"
	StatusInProgress Status = "In Progress"
	StatusDone       Status = "Done"
)

type Task struct {
	gorm.Model  `gorm:"embedded"`
	ID          uuid.UUID
	UserID      uuid.UUID
	TaskListID  uuid.UUID
	Title       string
	Description string
	Status      Status
}
