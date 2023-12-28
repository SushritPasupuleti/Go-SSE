// Handles loading and validation of environment variables
package env

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

// Config struct with environment variables
type Config struct {
	PORT        string
	JWT_SECRET  string
	ENVIRONMENT string
}

var DefaultConfig Config

// Load and validate environment variables
// `DefaultConfig` is now available to use
func Load() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error loading .env file")
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal().
			Err(errors.New("$PORT must be set")).
			Msg("$PORT must be set")
		os.Exit(1)
	}

	jwt_secret := os.Getenv("JWT_SECRET")
	if jwt_secret == "" {
		log.Fatal().
			Err(errors.New("$JWT_SECRET must be set")).
			Msg("$JWT_SECRET must be set")
		os.Exit(1)
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		log.Fatal().
			Err(errors.New("$ENVIRONMENT must be set")).
			Msg("$ENVIRONMENT must be set")
		os.Exit(1)
	}

	DefaultConfig = Config{
		PORT:        port,
		JWT_SECRET:  jwt_secret,
		ENVIRONMENT: environment,
	}

	// log.Info().Msgf("Successfully loaded environment variables: %v", DefaultConfig)
	log.Info().Msg("Successfully loaded environment variables")

	return DefaultConfig, nil
}
