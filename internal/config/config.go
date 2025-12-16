package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type DatabaseConfig struct {
	DatabaseURL string
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type JWTConfig struct {
	SecretKey       string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

// Init initializes configuration by parsing YAML file and environment variables, returns populated Config or error if parsing fails.
func Init() (*Config, error) {
	var cfg Config

	if err := parseYml(&cfg); err != nil {
		return nil, err
	}

	err := parseEnvs(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// parseEnvs parses env's from .env file and populates the provided Config struct, returns error if parsing fails.
func parseEnvs(cfg *Config) error {
	err := godotenv.Load("./.env")
	if err != nil {
		return fmt.Errorf("failed to find .env: %w", err)
	}

	cfg.Database.DatabaseURL = os.Getenv("DB_URL")
	if cfg.Database.DatabaseURL == "" {
		return fmt.Errorf("DB_URL is required")
	}

	cfg.JWT.SecretKey = os.Getenv("JWT_SECRET_KEY")
	if cfg.JWT.SecretKey == "" {
		return fmt.Errorf("JWT_SECRET_KEY is required")
	}

	return nil
}

// parseYml parses YAML config file and populates the provided Config struct, returns error if parsing fails.
func parseYml(cfg *Config) error {
	// Name of the config file without an extension (Viper will intuit the type
	// from an extension on the actual file)
	viper.SetConfigName("main")

	// Add search paths to find the file
	viper.AddConfigPath("./configs")

	// Find and read the config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("failed to unmarshal to struct: %w", err)
	}

	return nil
}
