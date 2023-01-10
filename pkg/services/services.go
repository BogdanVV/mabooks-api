package services

import (
	"github.com/bogdanvv/mabooks-api"
	"github.com/bogdanvv/mabooks-api/pkg/repository"
)

type Services struct {
	Authorization
}

type Authorization interface {
	SignUp(signUpBody models.SignUpBody) (string, error)
	Login(signInBody models.LoginBody) (models.LoginResponse, error)
	HandleToken(token string)
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Authorization: NewAuthorizationService(repo),
	}
}
