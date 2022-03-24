package core

import "github.com/go-pg/pg/v10/orm"

const RoleTable = "roles"

func init() {
	orm.RegisterTable((*RoleHasPermissions)(nil))
}

type Role struct {
	tableName   struct{} `sql:"roles"`
	ID          int
	Name        string
	DisplayName string
	Description string
	Permissions []Permission `pg:"many2many:role_has_permissions"`
}

type RoleHasPermissions struct {
	tableName    struct{} `sql:"role_has_permissions"`
	RoleId       int
	PermissionId int
}
