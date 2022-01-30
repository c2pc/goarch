package service

import (
	"github.com/chincharovpc/goarch/repository"
)

type UserService interface {
}

type User struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *User {
	return &User{userRepository: userRepository}
}
