package driven

import "tanakrit.assignment4.task/internal/core/domain"

type TaskCache interface {
	Save(task domain.Task) (bool, error)
	Update(task domain.Task) (bool, error)
	Delete(id uint) (bool, error)
}