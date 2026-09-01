package service

import (
	"myapp/internal/models"
	"myapp/internal/repository"
)

type TaskStatusesService struct {
	Repo *repository.TaskStatusesRepository
}

func (s *TaskStatusesService) GetTaskStatusByID(id int) (*models.TaskStatus, error) {
	return s.Repo.GetTaskStatusByID(id)
}

func (s *TaskStatusesService) GetAllTaskStatuses() ([]models.TaskStatus, error) {
	return s.Repo.GetAllTaskStatuses()
}
