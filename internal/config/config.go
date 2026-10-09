package config

import (
	"os"
	"time"
)

type Config struct {
	Port string
	DatabaseURL string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	IdleTimeout time.Duration
	JwtSecret string
}

func Load() Config {
	return Config{
		Port: getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://root:root@localhost:5432/shortr?sslmode=disable"),
		JwtSecret: getEnv("JWT_SECRET", "secret"),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 60 * time.Second,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}