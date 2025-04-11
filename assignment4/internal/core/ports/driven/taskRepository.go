package driven

import "tanakrit.assignment4.task/internal/core/domain"

type TaskRepository interface {
	Insert(task domain.Task) (domain.Task, error)
	Update(task domain.Task) (domain.Task , error)
	GetById(id uint) (domain.Task, error)
	RemoveById(id uint) (bool, error)
}