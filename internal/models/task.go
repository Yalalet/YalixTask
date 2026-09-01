package models

import "time"

type Task struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	CreatedAt   time.Time  `json:"created_at"`
	Deadline    time.Time  `json:"deadline"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	Description string     `json:"description"`
	TeamID      *int       `json:"team_id,omitempty"`
	StatusID    int        `json:"status_id"`
	PriorityID  int        `json:"priority_id"`
}

func (t *Task) IsCompleted() bool {
	return t.CompletedAt != nil
}

func (t *Task) IsOverdue() bool {
	if t.CompletedAt != nil {
		return false
	}
	return time.Now().After(t.Deadline)
}

//-------------safe struct ------------

type TaskPublic struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	Deadline    time.Time `json:"deadline"`
	Description string    `json:"description"`
	TeamID      *int      `json:"team_id,omitempty"`
	StatusID    int       `json:"status_id"`
	PriorityID  int       `json:"priority_id"`
}

func (t *Task) ToPublicTask() TaskPublic {
	return TaskPublic{
		ID:          t.ID,
		Name:        t.Name,
		CreatedAt:   t.CreatedAt,
		Deadline:    t.Deadline,
		Description: t.Description,
		TeamID:      t.TeamID,
		StatusID:    t.StatusID,
		PriorityID:  t.PriorityID,
	}
}
