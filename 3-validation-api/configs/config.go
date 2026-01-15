package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Mailer MailerConfig
}

type MailerConfig struct {
	Email string;
	Password string;
	Host string;
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file!")
		panic("Could not load .env file")
	}
	return &Config{
		Mailer: MailerConfig{
			Email: os.Getenv("MAILER_EMAIL"),
			Password: os.Getenv("MAILER_PASSWORD"),
			Host: os.Getenv("MAILER_HOST"),
		},
	}
}