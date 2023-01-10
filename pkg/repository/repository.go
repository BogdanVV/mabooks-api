package repository

import (
	"fmt"
	"log"

	models "github.com/bogdanvv/mabooks-api"
	"github.com/jmoiron/sqlx"
)

const (
	host        = "localhost"
	port        = "5430"
	user        = "postgres"
	password    = "qweqwe"
	dbname      = "postgres"
	users_table = "users"
)

type Authorization interface {
	CreateUser(signUpBody models.SignUpBody) (string, error)
	GetUserByLoginData(email string, hashedPassword string) (models.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
	}
}

func ConnectToDB() (*sqlx.DB, error) {
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sqlx.Open("postgres", psqlConn)
	if err != nil {
		log.Fatalf("Failed to open db: %s", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping-pong >>> %s", err.Error())
		return nil, err
	}

	return db, nil
}
