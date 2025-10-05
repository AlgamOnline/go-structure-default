package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string

	SocketHost string
	SocketPort int
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func LoadConfig() *Config {
	// Load .env jika ada
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))

	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	socketPort, err := strconv.Atoi(getEnv("SOCKET_PORT", "5050"))
	if err != nil {
		log.Fatalf("Invalid SOCKET_PORT: %v", err)
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     dbPort,
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "mydb"),

		SocketHost: getEnv("SOCKET_HOST", "localhost"),
		SocketPort: socketPort,
	}
}
