package repositories

import (
	"fmt"

	"tanakrit.karaket/assignment3/models"
)

type TaskInMemo struct {
	TasksMapper     map[int]models.Task
	IdRunningNumber int
}

func InitMemo() (taskMemo *TaskInMemo) {
	taskMemo = new(TaskInMemo)
	taskMemo.TasksMapper = make(map[int]models.Task)
	taskMemo.IdRunningNumber = 0
	return
}

func (taskMemo *TaskInMemo) Add(task models.Task) {
	taskMemo.IdRunningNumber++
	task.Id = taskMemo.IdRunningNumber
	taskMemo.TasksMapper[taskMemo.IdRunningNumber] = task
}

func (taskMemo *TaskInMemo) GetIdRunningNumber() int {
	return taskMemo.IdRunningNumber
}

func (taskMemo *TaskInMemo) GetTask(id int) models.Task {
	return taskMemo.TasksMapper[id]
}

func (taskMemo *TaskInMemo) GetTasks() []models.Task {
	tasks := make([]models.Task, taskMemo.IdRunningNumber)
	index := 0
	for _, task := range taskMemo.TasksMapper {
		fmt.Println(task)
		tasks[index] = task
		index += 1
	}
	return tasks
}

func (taskMemo *TaskInMemo) Replace(task models.Task) {
	taskMemo.TasksMapper[task.Id] = task
}

func (taskMemo *TaskInMemo) Remove(id int) {
	delete(taskMemo.TasksMapper, id)
}
