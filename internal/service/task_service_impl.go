package service

import (
	"context"

	pb "github.com/itmrchow/todolist-proto/protobuf/task"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"itmrchow/todolist-task/internal/repository"
)

type taskServiceImpl struct {
	pb.UnimplementedTaskServiceServer
	taskRepo repository.TasksRepository
}

func NewTaskService(taskRepo repository.TasksRepository) pb.TaskServiceServer {
	return &taskServiceImpl{
		taskRepo: taskRepo,
	}
}

func (s *taskServiceImpl) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}

func (s *taskServiceImpl) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}

func (s *taskServiceImpl) FindTask(ctx context.Context, req *pb.FindTaskRequest) (*pb.FindTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTask not implemented")
}

func (s *taskServiceImpl) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
