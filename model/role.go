package model

type Role struct {
	ID          int
	Name        string
	DisplayName string
	Description string
	Permissions []RoleHasPermissions `pg:"many2many:role_has_permissions"`
}

type RoleHasPermissions struct {
	RoleID       int
	PermissionID int
}
