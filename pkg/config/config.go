package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// Config holds all HTTP configuration.
const LogLevel = "INFO"
const Environment = "dev"

type HTTPConfig struct {
	Host       string
	Port       string
	ExposePort string
}

func LoadHTTPConfig() HTTPConfig {

	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "8100"
	}
	return HTTPConfig{
		Host:       os.Getenv("HOST"),
		Port:       appPort,
		ExposePort: os.Getenv("EXPOSE_PORT"),
	}
}

type AppConfig struct {
	HTTP        HTTPConfig
	Environment string         `json:"environment"`
	Validator   echo.Validator `json:"-"`
}

func NewConfig() (*AppConfig, error) {
	var err error
	currentEnv := "development"
	// load env file if exists
	_, err = os.Stat(".env")
	if err == nil {
		err = godotenv.Load(os.ExpandEnv(".env"))
		if err != nil {
			return nil, fmt.Errorf("error initializing app: %v", err)
		}
	}

	if os.Getenv("ENV") != "" {
		currentEnv = os.Getenv("ENV")
	}

	return &AppConfig{
		Environment: currentEnv,
		HTTP:        LoadHTTPConfig(),
		Validator:   &AppValidator{validator: validator.New()},
	}, nil
}
