package v1

import (
	"github.com/chincharovpc/goarch/controller"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	userController controller.UserController
}

func NewUserRouters(user controller.UserController) *UserRouter {
	return &UserRouter{
		userController: user,
	}
}

func (r *UserRouter) Init(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.GET("/:id", r.userController.GetByID)
	}
}
