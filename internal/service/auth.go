package service

type IAuthService interface {
}

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}
