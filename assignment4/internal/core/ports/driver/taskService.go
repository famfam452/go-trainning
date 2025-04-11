package driver

import (
	"tanakrit.assignment4.task/internal/core/domain/models/requests"
	"tanakrit.assignment4.task/internal/core/domain/models/responses"
)

type TaskService interface {
	AddTask(request requests.AddTaskRequest) (responses.AddTaskResponse, error)
	UpdateTask(id uint, request requests.UpdateTaskRequest) (responses.UpdateTaskResponse, error)
	DeleteTask(id uint) (responses.DeleteTaskResponse, error)
}