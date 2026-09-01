package repository

import (
	"database/sql"
	"myapp/internal/models"
)

type TeamRoleRepository struct {
	DB *sql.DB // Add any necessary fields, such as a database connection
}

func (r *TeamRoleRepository) GetTeamRoleByID(id int) (*models.TeamRole, error) {
	var teamrole models.TeamRole
	err := r.DB.QueryRow(`SELECT id, name, description, created_at, updated_at FROM team_roles WHERE id = $1`,
		id).Scan(&teamrole.ID, &teamrole.Name, &teamrole.Description, &teamrole.CreatedAt, &teamrole.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &teamrole, nil
}

func (r *TeamRoleRepository) GetAllTeamRoles() ([]models.TeamRole, error) {
	rows, err := r.DB.Query("SELECT id, name, description, created_at, updated_at FROM team_roles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teamroles []models.TeamRole
	for rows.Next() {
		var teamrole models.TeamRole
		err := rows.Scan(&teamrole.ID, &teamrole.Name, &teamrole.Description, &teamrole.CreatedAt, &teamrole.UpdatedAt)
		if err != nil {
			return nil, err
		}
		teamroles = append(teamroles, teamrole)
	}
	return teamroles, nil
}
