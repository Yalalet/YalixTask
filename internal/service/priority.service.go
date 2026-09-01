package service

import (
	"myapp/internal/models"
	"myapp/internal/repository"
)

type PriorityService struct {
	Repo *repository.PriorityRepository
}

func (s *PriorityService) GetAllPriority() ([]models.Priority, error) {
	return s.Repo.GetAllPriority()
}
