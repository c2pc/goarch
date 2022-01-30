package model

import "time"

type User struct {
	Id              int
	Username        string
	Name            string
	Password        *string
	Email           string
	Token           *string
	EmailVerifiedAt *time.Time
	UpdatedAt       time.Time
	CreatedAt       time.Time
	Roles           []UserHasRoles `pg:"many2many:user_has_roles"`
}

type UserHasRoles struct {
	UserID int
	RoleID int
}
