package service

import (
	"myapp/internal/models"
	"myapp/internal/repository"
)

type TeamUserService struct {
	Repo *repository.TeamUserRepository
}

func (s *TeamUserService) CreateTeamUser(teamUser *models.TeamUser) error {
	err := s.Repo.CreateTeamUser(teamUser)
	return err
}

func (s *TeamUserService) GetAllTeamUsers() ([]models.TeamUser, error) {
	teamUsers, err := s.Repo.GetAllTeamUsers()
	if err != nil {
		return nil, err
	}
	return teamUsers, nil
}
