package services

import (
	models "github.com/bogdanvv/mabooks-api"
	"github.com/bogdanvv/mabooks-api/pkg/repository"
)

type ReadBookServices struct {
	repo repository.ReadBookRepository
}

func NewReadBookService(repo repository.ReadBookRepository) *ReadBookServices {
	return &ReadBookServices{repo: repo}
}

func (s *ReadBookServices) CreateReadBook(userId string, readBook models.ReadBookInput) (string, error) {
	return s.repo.CreateReadBook(userId, readBook)
}

func (s *ReadBookServices) GetBookById(bookId string) (models.ReadBook, error) {
	return s.repo.GetBookById(bookId)
}

func (s *ReadBookServices) GetAllBooksByUserId(userId string) ([]models.ReadBook, error) {
	return s.repo.GetAllBooksByUserId(userId)
}

func (s *ReadBookServices) DeleteBook(bookId string) (string, error) {
	return s.repo.DeleteBook(bookId)
}

func (s *ReadBookServices) UpdateBook(bookId string, updateBody models.ReadBookInput) (models.ReadBook, error) {
	return s.repo.UpdateBook(bookId, updateBody)
}
