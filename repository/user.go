package repository

import (
	"github.com/chincharovpc/goarch/model"
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
)

type UserRepository interface {
}

type User struct {
	db     *pg.DB
	logger *zap.Logger
	model  *model.User
}

func NewUserRepository(db *pg.DB, logger *zap.Logger) *User {
	return &User{
		db:     db,
		logger: logger,
	}
}
