package service

import (
	"myapp/internal/models"
	"myapp/internal/repository"
)

type TeamService struct {
	Repo *repository.TeamRepository
}

func (s *TeamService) GetAllTeams() ([]models.Team, error) {
	return s.Repo.GetAllTeams()
}
