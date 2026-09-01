package service

import (
	"myapp/internal/models"
	"myapp/internal/repository"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func (s *TaskService) CreateTask(task *models.Task) error {
	err := s.Repo.CreateTask(task)
	return err
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.Repo.GetAllTasks()
}

func (s *TaskService) GetTaskByID(id int) (*models.Task, error) {
	return s.Repo.GetTaskByID(id)
}
