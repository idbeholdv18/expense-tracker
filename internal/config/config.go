package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host        string
	Port        string
	TLSFile     string
	TLSKey      string
	DatabaseURL string
	JWTSecret   string
	CORSOrigin  string
	BcryptCost  int
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	tlsFile := os.Getenv("TLS_FILE")
	if tlsFile == "" {
		tlsFile = "cert/cert.pem"
	}

	tlsKey := os.Getenv("TLS_KEY")
	if tlsKey == "" {
		host = "cert/key.pem"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	cors := os.Getenv("CORS_ORIGIN")
	if cors == "" {
		cors = "http://localhost:3000"
	}

	cost := 10
	if envCost := os.Getenv("BCRYPT_COST"); envCost != "" {
		if parsed, err := strconv.Atoi(envCost); err == nil {
			cost = parsed
		}
	}

	return &Config{
		Host:        host,
		Port:        port,
		TLSFile:     tlsFile,
		TLSKey:      tlsKey,
		DatabaseURL: dbURL,
		JWTSecret:   jwtSecret,
		CORSOrigin:  cors,
		BcryptCost:  cost,
	}
}
