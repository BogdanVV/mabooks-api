package services

import (
	"github.com/bogdanvv/mabooks-api"
	"github.com/bogdanvv/mabooks-api/pkg/repository"
)

type Services struct {
	Authorization
	ReadBook
}

type ReadBook interface {
	CreateReadBook(userId string, readBook models.ReadBookInput) (string, error)
	GetBookById(bookId string) (models.ReadBook, error)
	GetAllBooksByUserId(userId string) ([]models.ReadBook, error)
	UpdateBook(bookId string, updateBody models.ReadBookInput) (models.ReadBook, error)
	DeleteBook(bookId string) (string, error)
}

type Authorization interface {
	SignUp(signUpBody models.SignUpBody) (string, error)
	Login(models.User) (models.LoginResponse, error)
	HandleToken(token string)
	ReissueTokens(refreshToken string) (models.TokenPair, error)
	GetUserByEmail(email string) (models.User, error)
	VerifyPassword(password string, user models.User) error
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Authorization: NewAuthorizationService(repo),
		ReadBook:      NewReadBookService(repo),
	}
}
