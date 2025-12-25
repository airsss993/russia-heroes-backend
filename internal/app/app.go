package app

import (
	"github.com/airsss993/russia-heroes-backend/internal/config"
	"github.com/airsss993/russia-heroes-backend/pkg/logger"
	"github.com/airsss993/russia-heroes-backend/pkg/utils"
	"go.uber.org/zap"
)

// Run инициализирует и запускает приложение
func Run() {
	if err := logger.Init("dev"); err != nil {
		panic(err)
	}
	defer logger.L.Sync()

	_, err := config.Init()
	if err != nil {
		logger.L.Fatal("ошибка инициализации конфига", zap.Error(err))
	}

	logger.L.Info("запуск russia-heroes backend",
		zap.String("version", "0.1.0"),
	)

	creds := utils.GenerateAdminCredentials()
	utils.PrintCredentials(creds)
}
