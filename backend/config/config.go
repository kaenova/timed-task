package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	H HTTPConfig

	D DatabaseConfig
	R RabbitMQConfig
}

func MakeConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env is not laoded")
	}
	return Config{
		H: MakeHttpConfig(),

		D: MakeDatabaseConfig(),
		R: MakeRabbitMQConfig(),
	}
}

func GetEnvOrDefault(key string, fallback string) string {
	env := os.Getenv(key)
	if env == "" {
		env = fallback
	}
	return env
}
