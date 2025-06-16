package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("karhub.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados SQLite: %v", err)
	}

	DB = db
	log.Println("Conectado ao banco de dados SQLite com sucesso!")
}
