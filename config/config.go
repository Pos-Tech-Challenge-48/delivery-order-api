package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type postgresqlDBConfig struct {
	ConnectionString string
}

type Environment struct {
	Name string
}

var EnvLocal = Environment{Name: "local"}
var EnvTest = Environment{Name: "test"}
var EnvHomolog = Environment{Name: "homolog"}
var EnvSandbox = Environment{Name: "sandbox"}
var EnvProduction = Environment{Name: "production"}

type SQSConfig struct {
	LocalURL          string
	SessionMaxRetries int
	MaxWorkers        int
	MaxMessages       int
	PollInterval      time.Duration
	VisibilityTimeout time.Duration

	// QUEUES
	QueuePaymentsConfirmation string
	QueueOrderProduction      string
}

type Config struct {
	DBConfig  postgresqlDBConfig
	SQSConfig SQSConfig

	Environment Environment
}

func LoadConfig() (*Config, error) {

	environ := Environment{
		Name: os.Getenv("ENVIRONMENT"),
	}

	if environ.Name != EnvProduction.Name {
		if err := godotenv.Load(); err != nil {
			fmt.Printf("Error REASON: %v", err)
			return nil, errors.New("error to load env")
		}
	}

	pgsqlConfig := postgresqlDBConfig{ConnectionString: os.Getenv("POSTGRESQL_URL")}

	// SQS
	SessionMaxRetries, _ := strconv.Atoi(os.Getenv("SESSION_MAX_RETRIES"))
	MaxWorkers, _ := strconv.Atoi(os.Getenv("SQS_MAX_WORKERS"))
	MaxMessages, _ := strconv.Atoi(os.Getenv("SQS_MAX_MESSAGES"))

	PollInterval, _ := time.ParseDuration(os.Getenv("SQS_POLL_INTERVAL"))
	VisibilityTimeout, _ := time.ParseDuration(os.Getenv("SQS_VISIBILITY_TIMEOUT"))

	SQSConfig := SQSConfig{
		LocalURL:          os.Getenv("SQS_LOCAL_URL"),
		SessionMaxRetries: SessionMaxRetries,
		MaxWorkers:        MaxWorkers,
		MaxMessages:       MaxMessages,
		PollInterval:      PollInterval,
		VisibilityTimeout: VisibilityTimeout,

		QueuePaymentsConfirmation: os.Getenv("SQS_PAYMENTS_CONFIRMATION_QUEUE"),
		QueueOrderProduction:      os.Getenv("SQS_ORDER_PRODUCTION_QUEUE"),
	}

	config := Config{DBConfig: pgsqlConfig, SQSConfig: SQSConfig, Environment: environ}

	return &config, nil
}
