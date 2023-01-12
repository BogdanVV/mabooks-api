package repository

import (
	"fmt"

	models "github.com/bogdanvv/mabooks-api"
	"github.com/jmoiron/sqlx"
)

type ReadBookDB struct {
	db *sqlx.DB
}

func NewReadBookDB(db *sqlx.DB) *ReadBookDB {
	return &ReadBookDB{db: db}
}

func (db *ReadBookDB) CreateReadBook(userId string, readBook models.ReadBookInput) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (user_id, title, author, notes, is_finished) VALUES ($1, $2, $3, $4, $5) RETURNING id", read_books_table)
	result := db.db.QueryRow(query, userId, readBook.Title, readBook.Author, readBook.Notes, readBook.IsFinished)

	err := result.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (db *ReadBookDB) GetBookById(bookId string) (models.ReadBook, error) {
	var book models.ReadBook
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", read_books_table)
	err := db.db.Get(&book, query, bookId)

	return book, err
}

func (db *ReadBookDB) GetAllBooksByUserId(userId string) ([]models.ReadBook, error) {
	var books []models.ReadBook

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", read_books_table)
	err := db.db.Select(&books, query, userId)

	return books, err
}
