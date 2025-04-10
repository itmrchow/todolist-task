package repository

import (
	"context"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"itmrchow/todolist-task/infra"
	"itmrchow/todolist-task/internal/entity"
)

func TestTaskSuite(t *testing.T) {
	suite.Run(t, new(TaskTestSuite))
}

type TaskTestSuite struct {
	suite.Suite
	taskRepo TasksRepository
	db       *gorm.DB
}

func (s *TaskTestSuite) SetupTest() {
	db, _ := infra.InitSqlliteDb()

	sqlDB, _ := db.DB()

	// init test data
	fixtures, _ := testfixtures.New(
		testfixtures.Database(sqlDB),       // You database connection
		testfixtures.Dialect("sqlite"),     // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("testdata"), // The directory containing the YAML files
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)

	err := fixtures.Load()
	s.Require().NoError(err)

	s.taskRepo = NewTasksRepository(db)
	s.db = db
}

func (s *TaskTestSuite) TestCreateTask() {
	// prepare
	task := &entity.Task{
		ID:          uuid.New(),
		UserID:      uuid.New(),
		TaskListID:  nil,
		Title:       "Test Task",
		Description: "Test Description",
		Status:      entity.StatusPending,
	}

	// execute
	err := s.taskRepo.CreateTask(context.Background(), task)
	s.Assert().NoError(err)

	// assert
	var createdTask entity.Task
	s.db.First(&createdTask, "id = ?", task.ID)
	s.Assert().Equal(task.Title, createdTask.Title)
	s.Assert().Equal(task.Description, createdTask.Description)
	s.Assert().Equal(task.Status, createdTask.Status)
}

func (s *TaskTestSuite) TestGetTask_GetParentTask() {
	// prepare
	taskID := "123e4567-e89b-12d3-a456-426614174000"

	// execute
	createdTask, err := s.taskRepo.GetTask(context.Background(), taskID, PreloadOption{})
	s.Assert().NoError(err)

	// assert
	s.Assert().Equal(taskID, createdTask.ID.String())
	s.Assert().Equal("123e4567-e89b-12d3-a456-426614174001", createdTask.UserID.String())
	s.Assert().Nil(createdTask.TaskListID)
	s.Assert().Equal("test-title-1", createdTask.Title)
	s.Assert().Equal("test-description-1", createdTask.Description)
	s.Assert().Equal(entity.StatusPending, createdTask.Status)
}

func (s *TaskTestSuite) TestGetTask_GetParentTask_WithChildTask() {
	// prepare
	taskID := "123e4567-e89b-12d3-a456-426614174000"

	// execute
	createdTask, err := s.taskRepo.GetTask(context.Background(), taskID, PreloadOption{WithTaskList: true})
	s.Assert().NoError(err)

	// assert
	s.Assert().Equal(taskID, createdTask.ID.String())
	s.Assert().Equal("123e4567-e89b-12d3-a456-426614174001", createdTask.UserID.String())
	s.Assert().Nil(createdTask.TaskListID)
	s.Assert().Equal("test-title-1", createdTask.Title)
	s.Assert().Equal("test-description-1", createdTask.Description)
	s.Assert().Equal(entity.StatusPending, createdTask.Status)

	s.Assert().NotNil(createdTask.TaskList)
	s.Assert().Equal(1, len(*createdTask.TaskList))
	s.Assert().Equal("123e4567-e89b-12d3-a456-426614174001", (*createdTask.TaskList)[0].ID.String())
}

func (s *TaskTestSuite) TestGetTask_GetChildTask() {
	// prepare
	taskID := "123e4567-e89b-12d3-a456-426614174001"

	// execute
	createdTask, err := s.taskRepo.GetTask(context.Background(), taskID, PreloadOption{})
	s.Assert().NoError(err)

	// assert
	s.Assert().Equal(taskID, createdTask.ID.String())
	s.Assert().Equal("123e4567-e89b-12d3-a456-426614174001", createdTask.UserID.String())
	s.Assert().Equal("123e4567-e89b-12d3-a456-426614174000", createdTask.TaskListID.String())
	s.Assert().Equal("sub-test-title-2", createdTask.Title)
	s.Assert().Equal("sub-test-description-2", createdTask.Description)
	s.Assert().Equal(entity.StatusPending, createdTask.Status)

}
