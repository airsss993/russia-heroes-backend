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

// Init инициализирует весь конфиг с уже заполненными полями
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

// Загрузка env переменных в конфиг
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

// Парсинг значений с yml конфига в структуру
func parseYml(cfg *Config) error {
	// Указать имя файла конфига
	viper.SetConfigName("main")

	// Добавить путь для поиска конфига
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("failed to unmarshal to struct: %w", err)
	}

	return nil
}
