package models

import "time"

type TaskAssignee struct {
	TaskID     int       `json:"task_id"`
	UserID     int       `json:"user_id"`
	AssignedBy int       `json:"assigned_by"`
	AssignedAt time.Time `json:"assigned_at"`
	CreatedAt  time.Time `json:"created_at"`
}
