package repository

import (
	"database/sql"
	"fmt"
	"myapp/internal/models"
)

type PriorityRepository struct {
	DB *sql.DB // Add any necessary fields, such as a database connection
}

func (r *PriorityRepository) GetAllPriority() ([]models.Priority, error) {
	rows, err := r.DB.Query("SELECT id, name , color FROM priority")
	fmt.Println(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prioritys []models.Priority
	for rows.Next() {
		var priority models.Priority
		err := rows.Scan(&priority.ID, &priority.Name, &priority.Color)
		if err != nil {
			return nil, err
		}
		prioritys = append(prioritys, priority)
	}
	return prioritys, nil
}
