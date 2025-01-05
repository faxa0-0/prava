package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}
	return Config{
		DatabaseURL: Checkup("DATABASE_URL", "postgres://postgres:faxa@localhost:5432/tempi"),
		ServerPort:  Checkup("SERVER_PORT", "8887"),
	}, nil
}

func Checkup(env, placeholder string) string {
	res := os.Getenv(env)
	if res == "" {
		res = placeholder
	}
	return res
}
