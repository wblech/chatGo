package settings

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type GlobalConfig struct {
	KeycloakClientID     string
	KeycloakClientSecret string
	DbUsername           string
	DbPassword           string
	DbHost               string
}

func NewGlobalConfig() *GlobalConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return &GlobalConfig{
		KeycloakClientID:     os.Getenv("KEYCLOAK.CLIENTID"),
		KeycloakClientSecret: os.Getenv("KEYCLOAK.CLIENTSECRET"),
		DbUsername:           os.Getenv("DB.USERNAME"),
		DbPassword:           os.Getenv("DB.PASSWORD"),
		DbHost:               os.Getenv("DB.HOST"),
	}
}
