package repository

import (
	"database/sql"
	"myapp/internal/models"
)

type TaskRepository struct {
	DB *sql.DB // Add any necessary fields, such as a database connection
}

func (r *TaskRepository) GetTaskByID(id int) (*models.Task, error) {
	var task models.Task
	err := r.DB.QueryRow(`SELECT id, name, created_at, deadline, completed_at, description, team_id, status_id, priority_id 
			FROM tasks WHERE id = $1`,
		id).Scan(&task.ID, &task.Name, &task.CreatedAt, &task.Deadline, &task.CompletedAt, &task.Description, &task.TeamID, &task.StatusID, &task.PriorityID)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	err := r.DB.QueryRow(
		`INSERT INTO tasks (name, created_at, deadline, description, team_id, status_id, priority_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		task.Name, task.CreatedAt, task.Deadline, task.Description, task.TeamID, task.StatusID, task.PriorityID,
	).Scan(&task.ID)

	return err
}

func (r *TaskRepository) GetAllTasks() ([]models.Task, error) {
	rows, err := r.DB.Query("SELECT id, name, created_at, deadline, completed_at, description, team_id, status_id, priority_id FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Name, &task.CreatedAt, &task.Deadline, &task.CompletedAt, &task.Description, &task.TeamID, &task.StatusID, &task.PriorityID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
