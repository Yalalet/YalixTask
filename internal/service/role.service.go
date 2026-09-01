package service

import (
	"myapp/internal/models"
	"myapp/internal/repository"
)

type RoleService struct {
	Repo *repository.RoleRepository
}

func (s *RoleService) GetAllRoles() ([]models.Role, error) {
	return s.Repo.GetAllRoles()
}

func (s *RoleService) GetRoleByID(id int) (*models.Role, error) {
	return s.Repo.GetRoleByID(id)
}
