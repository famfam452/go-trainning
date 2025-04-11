package service

import (
	"tanakrit.assignment4.task/internal/core/domain"
	"tanakrit.assignment4.task/internal/core/domain/models/requests"
	"tanakrit.assignment4.task/internal/core/domain/models/responses"
	"tanakrit.assignment4.task/internal/core/ports/driven"
)

type TaskService struct {
	repository driven.TaskRepository
	cache driven.TaskCache
}

func NewTaskService(repository driven.TaskRepository, cache driven.TaskCache) *TaskService {
	return &TaskService{repository: repository, cache: cache}
}

func (service *TaskService) AddTask(request requests.AddTaskRequest) (responses.AddTaskResponse, error) {
	taskDomain := domain.Task{Title: request.Title, Description: request.Description, Completed: false}

	taskDomain, err  := service.repository.Insert(taskDomain)
	if err != nil {
		return responses.AddTaskResponse{}, err
	}

	_, err = service.cache.Save(taskDomain)
	if err != nil {
		return responses.AddTaskResponse{}, err
	}

	return responses.AddTaskResponse{ID: taskDomain.GetId(), Completed: taskDomain.GetCompleted()}, nil
}

func (service *TaskService) UpdateTask(id uint, request requests.UpdateTaskRequest) (responses.UpdateTaskResponse, error) {
	taskDomain := domain.Task{ID: id,Title: request.Title, Description: request.Description, Completed: request.Completed}

	taskDomain, err := service.repository.Update(taskDomain)
	if err != nil {
		return responses.UpdateTaskResponse{}, err
	}

	_, err = service.cache.Update(taskDomain)
	if err != nil {
		return responses.UpdateTaskResponse{}, err
	}

	return responses.UpdateTaskResponse{ID: taskDomain.GetId(), Title: taskDomain.GetTitle(), Description: taskDomain.GetDescription(),Completed: taskDomain.GetCompleted()}, nil
}

func (service *TaskService) DeleteTask(id uint) (responses.DeleteTaskResponse, error) {
	_, cachedError := service.cache.Delete(id)
	if cachedError != nil {
		return responses.DeleteTaskResponse{}, cachedError
	}
	_, repositoryError := service.repository.RemoveById(id)
	if repositoryError != nil {
		return responses.DeleteTaskResponse{}, repositoryError
	}

	return responses.DeleteTaskResponse{ID: id}, nil
}