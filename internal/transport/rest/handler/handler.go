package handler

import (
	"github.com/chincharovpc/goarch/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Handler struct {
	services *service.Services
}

func NewHandlers(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() *gin.Engine {
	handler := gin.Default()

	handler.Use(
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

	// Init handler
	handler.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(handler)

	return handler
}

func (h *Handler) initAPI(handler *gin.Engine) {
	api := handler.Group("/api")
	{
		NewAuthHandlers(h.services).Init(api)
	}

}
