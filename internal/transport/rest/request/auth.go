package request

import (
	"github.com/chincharovpc/goarch/pkg/apperr"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthSignUpRequest struct {
	Name            string `json:"name" binding:"required,min=3"`
	Email           string `json:"email" binding:"required,min=3,email"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

func AuthSignUp(c *gin.Context) (*AuthSignUpRequest, error) {
	var r AuthSignUpRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		return nil, err
	}
	if r.Password != r.PasswordConfirm {
		return nil, apperr.New(http.StatusBadRequest, "passwords do not match")
	}
	return &r, nil
}
