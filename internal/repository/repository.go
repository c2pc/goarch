package repository

import (
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
)

type Repositories struct {
	Role IRoleRepository
}

func NewRepositories(db *pg.DB, logger *zap.Logger) *Repositories {
	return &Repositories{
		Role: NewRoleRepository(db, logger),
	}
}
