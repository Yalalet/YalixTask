package repository

import (
	"database/sql"
	"myapp/internal/models"
)

type TaskAssigneeRepository struct {
	DB *sql.DB // Add any necessary fields, such as a database connection
}

func (r *TaskAssigneeRepository) GetTaskAssignee(taskID, userID int) (*models.TaskAssignee, error) {
	var assignee models.TaskAssignee
	err := r.DB.QueryRow(`SELECT task_id, user_id, assigned_by, assigned_at, created_at 
			FROM task_assignees WHERE task_id = $1 AND user_id = $2`,
		taskID, userID).Scan(&assignee.TaskID, &assignee.UserID, &assignee.AssignedBy, &assignee.AssignedAt, &assignee.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &assignee, nil
}

func (r *TaskAssigneeRepository) CreateTaskAssignee(taskAssignee *models.TaskAssignee) error {
	err := r.DB.QueryRow(
		`INSERT INTO task_assignees (task_id, user_id, assigned_by, created_at) 
		VALUES ($1, $2, $3, $4) RETURNING task_id`,
		taskAssignee.TaskID, taskAssignee.UserID, taskAssignee.AssignedBy, taskAssignee.CreatedAt,
	).Scan(&taskAssignee.TaskID)

	return err
}

func (r *TaskAssigneeRepository) GetAllTaskAssignees() (*sql.Rows, error) {
	rows, err := r.DB.Query(`SELECT task_id, user_id, assigned_by, assigned_at, created_at FROM task_assignees`)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
