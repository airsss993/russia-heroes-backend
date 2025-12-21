package app

import (
	"fmt"

	"github.com/airsss993/russia-heroes-backend/internal/config"
	"github.com/airsss993/russia-heroes-backend/pkg/logger"
	"go.uber.org/zap"
)

// Run инициализирует и запускает приложение
func Run() {
	if err := logger.Init("dev"); err != nil {
		panic(err)
	}
	defer logger.L.Sync()

	cfg, err := config.Init()
	if err != nil {
		logger.L.Fatal("failed to init config", zap.Error(err))
	}
	fmt.Println(cfg.Database.DatabaseURL)
	fmt.Println(cfg.Server.Port)

	logger.L.Info("starting russia-heroes backend",
		zap.String("version", "0.1.0"),
	)
}
