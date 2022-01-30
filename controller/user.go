package controller

import (
	"github.com/chincharovpc/goarch/service"
)

type UserController interface {
}

type User struct {
	service service.UserService
}

func NewUserController(user service.UserService) *User {
	return &User{service: user}
}
