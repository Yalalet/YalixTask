package models

import "time"

type TeamUser struct {
	TeamID     int        `json:"team_id"`
	UserID     int        `json:"user_id"`
	TeamRoleID int        `json:"team_role_id"`
	JoinedAt   time.Time  `json:"joined_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	LeftAt     *time.Time `json:"left_at"`
	IsActive   bool       `json:"is_active"`
	InvitedBy  int        `json:"invited_by"`
}
