package postgres

import (
	"gorm.io/gorm"
	"tanakrit.assignment4.task/internal/core/domain"
)

type TaskPostgresRepository struct {
	db *gorm.DB
}

func NewTaskPostgresRepository(db *gorm.DB) *TaskPostgresRepository{
	return &TaskPostgresRepository{db}
}
func (repository *TaskPostgresRepository) Insert(task domain.Task) (domain.Task, error) {
	taskEntity := Task{
		Title: task.GetTitle(),
		Description: task.GetDescription(),
		Completed: task.GetCompleted(),
	}
	err := repository.db.Create(&taskEntity).Error
	if err != nil {
		return task, err
	}
	return domain.Task{ID: taskEntity.ID, Title: taskEntity.Title, Description: taskEntity.Description, Completed: taskEntity.Completed}, nil
}
func (repository *TaskPostgresRepository) Update(task domain.Task) (domain.Task , error) {
	taskEntity := Task{
		ID: task.GetId(),
		Title: task.GetTitle(),
		Description: task.GetDescription(),
		Completed: task.GetCompleted(),
	}
	err := repository.db.Updates(&taskEntity).Error
	if err != nil {
		return task, err
	}
	return domain.Task{ID: taskEntity.ID, Title: taskEntity.Title, Description: taskEntity.Description, Completed: taskEntity.Completed}, nil
}
func (repository *TaskPostgresRepository) GetById(id uint) (domain.Task, error) {
	taskEntity := Task{ID: id}
	err := repository.db.Find(&taskEntity).Error
	if err != nil {
		return domain.Task{ID: id}, err
	}
	return domain.Task{ID: taskEntity.ID, Title: taskEntity.Title, Description: taskEntity.Description, Completed: taskEntity.Completed}, nil
}
func (repository *TaskPostgresRepository) RemoveById(id uint) (bool, error) {
	taskEntity := Task{ID: id}
	err := repository.db.Delete(&taskEntity).Error
	if err != nil {
		return false, err
	}
	return true, nil
}