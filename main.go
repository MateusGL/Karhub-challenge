package main

import (
	"karhub-beer-api/internal/config"
	"karhub-beer-api/internal/database"
	"karhub-beer-api/internal/models"
	"karhub-beer-api/internal/routes"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	database.Connect()
	database.DB.AutoMigrate(&models.Beer{})
	database.SeedBeers()

	r := routes.RegisterRoutes()

	port := config.GetPort()
	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
