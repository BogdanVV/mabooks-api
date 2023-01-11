package main

import (
	"log"

	// "github.com/bogdanvv/mabooks-api/pkg/examples"
	"github.com/bogdanvv/mabooks-api/pkg/handlers"
	"github.com/bogdanvv/mabooks-api/pkg/repository"
	"github.com/bogdanvv/mabooks-api/pkg/services"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")
	}

	db, err := repository.ConnectToDB()
	if err != nil {
		log.Fatal("Failed to connect to db")
	}

	repository := repository.NewRepository(db)
	services := services.NewServices(repository)
	handlers := handlers.NewHandlers(services)

	router := handlers.InitRoutes()

	/*
		examples.ExternalCallToAnotherAPI()
	*/

	router.Run(":9999")
}
