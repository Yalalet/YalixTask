package repository

import (
	"database/sql"
	"myapp/internal/models"
)

type TaskStatusesRepository struct {
	DB *sql.DB // Add any necessary fields, such as a database connection
}

func (r *TaskStatusesRepository) GetTaskStatusByID(id int) (*models.TaskStatus, error) {
	var taskStatus models.TaskStatus
	err := r.DB.QueryRow(`SELECT id, name, description, created_at, updated_at FROM task_statuses WHERE id = $1`,
		id).Scan(&taskStatus.ID, &taskStatus.Name, &taskStatus.Description, &taskStatus.CreatedAt, &taskStatus.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &taskStatus, nil
}

func (r *TaskStatusesRepository) GetAllTaskStatuses() ([]models.TaskStatus, error) {
	rows, err := r.DB.Query("SELECT id, name, description, created_at, updated_at FROM task_statuses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var taskStatuses []models.TaskStatus
	for rows.Next() {
		var taskStatus models.TaskStatus
		err := rows.Scan(&taskStatus.ID, &taskStatus.Name, &taskStatus.Description, &taskStatus.CreatedAt, &taskStatus.UpdatedAt)
		if err != nil {
			return nil, err
		}
		taskStatuses = append(taskStatuses, taskStatus)
	}
	return taskStatuses, nil
}
