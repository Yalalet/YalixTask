package repository

import (
	"database/sql"
	"myapp/internal/models"
)

type RoleRepository struct {
	DB *sql.DB // Add any necessary fields, such as a database connection
}

func (r *RoleRepository) GetRoleByID(id int) (*models.Role, error) {
	var role models.Role
	err := r.DB.QueryRow(`SELECT id , name FROM roles WHERE id = $1`,
		id).Scan(&role.ID, &role.Name)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) GetAllRoles() ([]models.Role, error) {
	rows, err := r.DB.Query("SELECT id, name FROM roles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		err := rows.Scan(&role.ID, &role.Name)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}
