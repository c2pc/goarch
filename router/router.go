package router

import (
	"github.com/chincharovpc/goarch/controller"
	"github.com/chincharovpc/goarch/router/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Router struct {
	controllers *controller.Controllers
}

func NewRouters(controllers *controller.Controllers) *Router {
	return &Router{
		controllers: controllers,
	}
}

func (r *Router) Init() *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return true
			},
			MaxAge: 12 * time.Hour,
		}),
	)

	// Init router
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.initAPI(router)

	return router
}

func (r *Router) initAPI(router *gin.Engine) {
	v1Routers := v1.NewRouters(r.controllers)
	api := router.Group("/api")
	{
		v1Routers.Init(api)
	}

}
