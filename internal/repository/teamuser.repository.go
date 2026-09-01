package repository

import (
	"database/sql"
	"myapp/internal/models"
)

type TeamUserRepository struct {
	DB *sql.DB // Add any necessary fields, such as a database connection
}

func (r *TeamUserRepository) GetTeamUser(teamID, userID int) (*models.TeamUser, error) {
	var teamUser models.TeamUser
	err := r.DB.QueryRow(`SELECT team_id, user_id, team_role_id, joined_at, created_at, updated_at, left_at, is_active, invited_by 
			FROM team_user WHERE team_id = $1 AND user_id = $2`,
		teamID, userID).Scan(&teamUser.TeamID, &teamUser.UserID, &teamUser.TeamRoleID, &teamUser.JoinedAt, &teamUser.CreatedAt, &teamUser.UpdatedAt, &teamUser.LeftAt, &teamUser.IsActive, &teamUser.InvitedBy)
	if err != nil {
		return nil, err
	}
	return &teamUser, nil
}

func (r *TeamUserRepository) CreateTeamUser(teamUser *models.TeamUser) error {
	err := r.DB.QueryRow(
		`INSERT INTO team_user (team_id, user_id, team_role_id, joined_at, created_at, invited_by) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING team_id`,
		teamUser.TeamID, teamUser.UserID, teamUser.TeamRoleID, teamUser.JoinedAt, teamUser.CreatedAt, teamUser.InvitedBy,
	).Scan(&teamUser.TeamID)

	return err
}

func (r *TeamUserRepository) GetAllTeamUsers() ([]models.TeamUser, error) {
	rows, err := r.DB.Query("SELECT team_id, user_id, team_role_id, joined_at, created_at, updated_at, left_at, is_active, invited_by FROM team_user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teamUsers []models.TeamUser
	for rows.Next() {
		var teamUser models.TeamUser
		if err := rows.Scan(&teamUser.TeamID, &teamUser.UserID, &teamUser.TeamRoleID, &teamUser.JoinedAt, &teamUser.CreatedAt, &teamUser.UpdatedAt, &teamUser.LeftAt, &teamUser.IsActive, &teamUser.InvitedBy); err != nil {
			return nil, err
		}
		teamUsers = append(teamUsers, teamUser)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return teamUsers, nil
}
