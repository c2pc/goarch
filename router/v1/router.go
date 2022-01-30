package v1

import (
	"github.com/chincharovpc/goarch/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
	controllers *controller.Controllers
}

func NewRouters(controllers *controller.Controllers) *Router {
	return &Router{
		controllers: controllers,
	}
}

func (r *Router) Init(api *gin.RouterGroup) {
	userRouters := NewUserRouters(r.controllers.User)

	v1 := api.Group("/v1")
	{
		userRouters.Init(v1)
	}
}
