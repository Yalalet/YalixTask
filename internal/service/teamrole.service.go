package service

import (
	"myapp/internal/models"
	"myapp/internal/repository"
)

type TeamRoleService struct {
	Repo *repository.TeamRoleRepository
}

func (s *TeamRoleService) GetAllTeamRoles() ([]models.TeamRole, error) {
	return s.Repo.GetAllTeamRoles()
}

func (s *TeamRoleService) GetTeamRoleByID(id int) (*models.TeamRole, error) {
	return s.Repo.GetTeamRoleByID(id)
}
