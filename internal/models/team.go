package models

import (
	"time"
)

type Team struct {
	ID          int       `json:"id"`
	Name        string    `json:"first_name"`
	CreatedAt   time.Time `json:"created_at"`
	Company     string    `json:"company"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
	Slug        string    `json:"slug"`
	PublicIs    bool      `json:"is_public"`
}

func (t Team) IsPublic() bool {
	return t.PublicIs != false
}
