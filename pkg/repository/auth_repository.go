package repository

import (
	"fmt"

	models "github.com/bogdanvv/mabooks-api"
	"github.com/jmoiron/sqlx"
)

type AuthDB struct {
	db *sqlx.DB
}

func NewAuthDB(db *sqlx.DB) *AuthDB {
	return &AuthDB{db: db}
}

func (db *AuthDB) CreateUser(signUpBody models.SignUpBody) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (email, password, username, phone, role) VALUES ($1, $2, $3, $4, $5) RETURNING id", users_table)
	response := db.db.QueryRow(query, signUpBody.Email, signUpBody.Password, signUpBody.Username, signUpBody.Phone, signUpBody.Role)
	err := response.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (a *AuthDB) GetUserByLoginData(email string, hashedPassword string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1 AND password = $2", users_table)
	err := a.db.Get(&user, query, email, hashedPassword)

	return user, err
}
