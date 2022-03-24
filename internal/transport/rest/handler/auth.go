package handler

import (
	"github.com/chincharovpc/goarch/internal/service"
	"github.com/chincharovpc/goarch/internal/transport/rest/request"
	"github.com/chincharovpc/goarch/pkg/apperr"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	services *service.Services
}

func NewAuthHandlers(services *service.Services) *AuthHandler {
	return &AuthHandler{services}
}

func (h *AuthHandler) Init(api *gin.RouterGroup) {
	api.POST("/signup", h.SignUp)
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	_, err := request.AuthSignUp(c)
	if err != nil {
		apperr.Response(c, err)
		return
	}

	c.Status(http.StatusCreated)
}
