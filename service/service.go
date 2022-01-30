package service

import (
	"github.com/chincharovpc/goarch/repository"
)

type Services struct {
	User UserService
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{
		User: NewUserService(repositories.User),
	}
}
