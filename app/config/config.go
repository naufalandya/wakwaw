package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type OracleConfig struct {
	Host     string
	Port     string
	Service  string
	User     string
	Password string
	Role     string
}

func LoadConfig() OracleConfig {
	_ = godotenv.Load()

	cfg := OracleConfig{
		Host:     os.Getenv("ORACLE_HOST"),
		Port:     os.Getenv("ORACLE_PORT"),
		Service:  os.Getenv("ORACLE_SERVICE"),
		User:     os.Getenv("ORACLE_USER"),
		Password: os.Getenv("ORACLE_PASSWORD"),
		Role:     os.Getenv("ORACLE_ROLE"),
	}

	fmt.Println("woii", cfg.Host)

	return cfg
}

func MustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("Missing required env: %s", key)
	}
	return v
}
