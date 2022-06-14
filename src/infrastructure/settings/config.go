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
	RMQUsername          string
	RMQPassword          string
	RMQHost              string
	RMQPort              string
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
		RMQUsername:          os.Getenv("RMQ.USERNAME"),
		RMQPassword:          os.Getenv("RMQ.PASSWORD"),
		RMQHost:              os.Getenv("RMQ.HOST"),
		RMQPort:              os.Getenv("RMQ.PORT"),
	}
}
