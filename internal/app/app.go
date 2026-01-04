package app

import (
	"github.com/airsss993/russia-heroes-backend/pkg/logger"
	"github.com/airsss993/russia-heroes-backend/pkg/utils"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

// Run инициализирует и запускает приложение
func Run() {
	if err := logger.Init("dev"); err != nil {
		panic(err)
	}
	defer logger.L.Sync()
	//
	//cfg, err := config.Init()
	//if err != nil {
	//	logger.L.Fatal("failed to init config", zap.Error(err))
	//}

	logger.L.Info("Starting russia-heroes backend",
		zap.String("version", "0.1.0"),
	)

	//ctx := context.Background()
	//db, err := sql.Open("postgres", cfg.Database.DatabaseURL)
	//if err != nil {
	//	logger.L.Fatal("failed to connect to db", zap.Error(err))
	//}
	//queries := sqlc.New(db)
	//adminRepo := repository.NewAdminRepo(queries)
	//adminService := service.NewAdminService(adminRepo, logger.L)
	//services := service.NewServices(adminService)
	//
	//err = services.AdminService.CreateSuperAdmin(ctx)
	creds := utils.GenerateAdminCredentials()
	utils.PrintCredentials(creds)
	//if err != nil {
	//	logger.L.Error("failed to create super-admin", zap.Error(err))
	//}
}
