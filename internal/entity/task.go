package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending"
	StatusInProgress TaskStatus = "in_progress"
	StatusDone       TaskStatus = "done"
)

type Task struct {
	gorm.Model  `gorm:"embedded"`
	ID          uuid.UUID  `gorm:"primaryKey"`
	UserID      uuid.UUID  `gorm:"not null"`
	TaskListID  *uuid.UUID `gorm:"index:idx_task_parent;foreignKey:ID;references:tasks"`
	Title       string     `gorm:"size:255;not null"`
	Description string     `gorm:"size:255;not null"`
	Status      TaskStatus `gorm:"size:32;not null"`
	TaskList    *[]Task    `gorm:"foreignKey:TaskListID;references:ID"`
}
