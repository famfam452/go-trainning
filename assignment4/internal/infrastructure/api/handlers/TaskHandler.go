package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
	"tanakrit.assignment4.task/internal/core/domain/models/requests"
	"tanakrit.assignment4.task/internal/core/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service}
}

func (handler *TaskHandler) PostAddTask(context echo.Context) error {
	var addTaskRequest requests.AddTaskRequest
	if err := context.Bind(&addTaskRequest); err != nil {
		log.Fatal(err)
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	response, err := handler.service.AddTask(addTaskRequest)
	if err != nil {
		log.Fatal(err)
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Request is not success"})
	}
	return context.JSON(http.StatusCreated, response)
}

func (handler *TaskHandler) PutUpdateTask(context echo.Context) error {
	var updateTaskRequest requests.UpdateTaskRequest
	id := context.Param("id")
	uintId, parseerr := strconv.ParseUint(id,10,64)
	if parseerr != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	if err := context.Bind(&updateTaskRequest); err != nil {
		log.Fatal(err)
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	response, err := handler.service.UpdateTask(uint(uintId), updateTaskRequest)
	if err != nil {
		log.Fatal(err)
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Request is not success"})
	}
	return context.JSON(http.StatusOK, response)
}

func (handler *TaskHandler) DeleteTaskById(context echo.Context) error {
	id := context.Param("id")
	uintId, parseerr := strconv.ParseUint(id,10,64)
	if parseerr != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	response, err := handler.service.DeleteTask(uint(uintId))
	if err != nil {
		log.Fatal(err)
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Request is not success"})
	}
	return context.JSON(http.StatusOK, response)
}