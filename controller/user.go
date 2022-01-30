package controller

import (
	"github.com/chincharovpc/goarch/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetByID(g *gin.Context)
}

type User struct {
	service service.UserService
}

func NewUserController(user service.UserService) *User {
	return &User{service: user}
}

func (c *User) GetByID(g *gin.Context) {

}
