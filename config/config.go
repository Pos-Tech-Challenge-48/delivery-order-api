package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type postgresqlDBConfig struct {
	ConnectionString string
}

type Config struct {
	DBConfig postgresqlDBConfig
}

func LoadConfig() (*Config, error) {
	if os.Getenv("ENVIRONMENT") != "production" {
		if err := godotenv.Load(); err != nil {
			return nil, errors.New("error to load env")
		}
	}

	pgsqlConfig := postgresqlDBConfig{ConnectionString: os.Getenv("POSTGRESQL_URL")}
	fmt.Printf("ENV do banco %s", pgsqlConfig)

	config := Config{DBConfig: pgsqlConfig}

	return &config, nil
}
