package config

import (
	"log"
	"os"
	"strings"
	
	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
	Cors CorsConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

type CorsConfig struct {
	AllowedOrigins   string
	AllowCredentials bool
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
		Cors: CorsConfig{
			AllowedOrigins: os.Getenv("CORS_ALLOWED_ORIGINS"),
			AllowCredentials: strings.Contains(
				"true,yes,1",
				strings.ToLower(os.Getenv("CORS_ALLOW_CREDENTIALS")),
			),
		},
	}
}
