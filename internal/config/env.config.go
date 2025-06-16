package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var requiredVars = []string{
	"PORT",
	"SPOTIFY_CLIENT_ID",
	"SPOTIFY_CLIENT_SECRET",
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env não encontrado, usando variáveis de ambiente do sistema")
	}

	missing := []string{}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			missing = append(missing, v)
		}
	}

	if len(missing) > 0 {
		log.Fatalf("Variáveis de ambiente ausentes: %s", strings.Join(missing, ", "))
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

func GetSpotifyCredentials() (string, string) {
	return os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_CLIENT_SECRET")
}
