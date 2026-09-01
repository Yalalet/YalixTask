package repository

import (
	"database/sql"
	"myapp/internal/models"
)

type TeamRepository struct {
	DB *sql.DB
}

func (r *TeamRepository) GetAllTeams() ([]models.Team, error) {
	rows, err := r.DB.Query("SELECT id , name , created_at , company , description , updated_at , slug , is_public FROM teams")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		err := rows.Scan(&team.ID, &team.Name, &team.CreatedAt, &team.Company, &team.Description, &team.UpdatedAt, &team.Slug, &team.PublicIs)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	return teams, nil
}
