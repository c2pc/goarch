package model

type Role struct {
	ID          int
	Name        string
	DisplayName string
	Description string
}

type RoleHasPermissions struct {
	RoleID       int
	PermissionID int
}
