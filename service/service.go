package service

import (
	"github.com/chincharovpc/goarch/repository"
)

type Services struct {
	User UserService
	Role RoleService
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{
		User: NewUserService(repositories.User),
		Role: NewRoleService(repositories.Role),
	}
}
