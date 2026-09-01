package service

import (
	"myapp/internal/models"
	"myapp/internal/repository"
)

type TaskAssigneeService struct {
	Repo *repository.TaskAssigneeRepository
}

func (s *TaskAssigneeService) CreateTaskAssignee(taskAssignee *models.TaskAssignee) error {
	err := s.Repo.CreateTaskAssignee(taskAssignee)
	return err
}

func (s *TaskAssigneeService) GetAllTaskAssignees() ([]models.TaskAssignee, error) {
	rows, err := s.Repo.GetAllTaskAssignees()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assignees []models.TaskAssignee
	for rows.Next() {
		var assignee models.TaskAssignee
		if err := rows.Scan(&assignee.TaskID, &assignee.UserID, &assignee.AssignedBy, &assignee.AssignedAt, &assignee.CreatedAt); err != nil {
			return nil, err
		}
		assignees = append(assignees, assignee)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return assignees, nil
}
