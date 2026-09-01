package models

import "time"

type User struct {
	ID           int        `json:"id"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	MiddleName   *string    `json:"middle_name,omitempty"`
	Login        string     `json:"login"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"-"`
	Avatar       *string    `json:"avatar,omitempty"`
	Phone        *string    `json:"phone,omitempty"`
	GitHub       *string    `json:"github,omitempty"`
	RoleName     string     `json:"role_name"`
	RoleID       int        `json:"role_id"`
	IsActive     bool       `json:"is_active"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
	LastLoginAt  *time.Time `json:"last_login_at,omitempty"`
}

func (u *User) IsDeleted() bool {
	return u.DeletedAt != nil
}

func (u *User) FullName() string {
	if u.MiddleName != nil && *u.MiddleName != "" {
		return u.FirstName + " " + *u.MiddleName + " " + u.LastName
	}

	return u.FirstName + " " + u.LastName
}

//-------------safe struct ------------

type UserPublic struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Avatar    *string   `json:"avatar,omitempty"`
	RoleID    int       `json:"role_id"`
	RoleName  string    `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) ToPublicUser() UserPublic {

	return UserPublic{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Avatar:    u.Avatar,
		RoleID:    u.RoleID,
		RoleName:  u.RoleName,
		CreatedAt: u.CreatedAt,
	}

}
