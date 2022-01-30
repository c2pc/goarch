package controller

import (
	"github.com/chincharovpc/goarch/service"
)

type Controllers struct {
	User UserController
}

func NewControllers(services *service.Services) *Controllers {
	return &Controllers{
		User: NewUserController(services.User),
	}
}
