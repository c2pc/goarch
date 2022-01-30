package repository

import (
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
)

type Repositories struct {
	User UserRepository
}

func NewRepositories(db *pg.DB, logger *zap.Logger) *Repositories {
	return &Repositories{
		User: NewUserRepository(db, logger),
	}
}
