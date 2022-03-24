package service

import (
	"github.com/chincharovpc/goarch/internal/repository"
)

type Services struct {
	Role IRoleService
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{
		Role: NewRoleService(repositories.Role),
	}
}
