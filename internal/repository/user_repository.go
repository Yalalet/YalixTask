package repository

import (
	"database/sql"
	"myapp/internal/models"
)

type UserRepository struct {
	DB *sql.DB // Add any necessary fields, such as a database connection
}

func (r *UserRepository) GetUserByLogin(login string) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow(`SELECT id, last_name , first_name, login, password_hash FROM users WHERE login = $1`,
		login).Scan(&user.ID, &user.LastName, &user.FirstName, &user.Login, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	err := r.DB.QueryRow(
		`INSERT INTO users (first_name , last_name , login , email, password_hash, role_id) 
		VALUES ($1, $2, $3, $4 , $5 , $6) RETURNING id`,
		user.FirstName, user.LastName, user.Login, user.Email, user.PasswordHash, user.RoleID,
	).Scan(&user.ID)

	return err
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.DB.Query(`
		SELECT u.id, u.last_name, u.first_name, u.email, u.login, r.name AS role_name , u.role_id
		FROM users u
		LEFT JOIN roles r ON u.role_id = r.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.LastName, &user.FirstName, &user.Email, &user.Login, &user.RoleName, &user.RoleID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

		if err := rows.Err(); err != nil {
			return nil, err
		}
	}
	return users, nil
}
