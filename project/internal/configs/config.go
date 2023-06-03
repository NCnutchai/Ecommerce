package configs

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}
	return nil
}
