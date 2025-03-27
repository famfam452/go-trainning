package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"tanakrit.karaket/assignment3/models"
	"tanakrit.karaket/assignment3/repositories"
)

var taskMemo = repositories.InitMemo()

func CreateTask(context echo.Context) error {
	task := new(models.Task)
	if err := context.Bind(task); err != nil {
		log.Fatal(err)
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	taskMemo.Add(*task)
	task.Id = taskMemo.GetIdRunningNumber()
	return context.JSON(http.StatusCreated, task)
}

func GetAllTasks(context echo.Context) error {
	tasks := taskMemo.GetTasks()
	return context.JSON(http.StatusCreated, tasks)
}

func GetTaskById(context echo.Context) error {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	task := taskMemo.GetTask(idInt)

	return context.JSON(http.StatusCreated, task)
}

func PatchTaskById(context echo.Context) error {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	task := new(models.Task)
	if err := context.Bind(task); err != nil {
		log.Fatal(err)
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	task.Id = idInt
	curTask := taskMemo.GetTask(task.Id)
	if task.Status != "" {
		curTask.Status = task.Status
	}
	if task.Title != "" {
		curTask.Title = task.Title
	}
	taskMemo.Replace(curTask)
	return context.JSON(http.StatusCreated, curTask)
}

func DeleteTask(context echo.Context) error {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	taskMemo.Remove(idInt)
	return context.JSON(http.StatusCreated, map[string]string{"message": "Task deleted successfully"})
}
