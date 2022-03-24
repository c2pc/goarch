package core

import "time"

type User struct {
	Id              int
	Email           string
	Name            string
	Password        *string
	Token           *string
	EmailVerifiedAt *time.Time
	UpdatedAt       time.Time
	CreatedAt       time.Time
	Roles           []Role `pg:"many2many:user_has_roles"`
}

type UserHasRoles struct {
	UserID int
	RoleID int
}
