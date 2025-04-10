package infra

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"itmrchow/todolist-task/internal/entity"
)

func InitSqlliteDb() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.Task{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	return db, nil
}
