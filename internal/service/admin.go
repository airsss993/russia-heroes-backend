package service

import (
	"context"
	"database/sql"

	"github.com/airsss993/russia-heroes-backend/internal/repository"
	"github.com/airsss993/russia-heroes-backend/internal/repository/postgres/sqlc"
	"github.com/airsss993/russia-heroes-backend/pkg/utils"
	"go.uber.org/zap"
)

type AdminService struct {
	repo repository.AdminRepo
	log  *zap.Logger
}

func NewAdminService(repo repository.AdminRepo, logger *zap.Logger) *AdminService {
	return &AdminService{
		repo: repo,
		log:  logger,
	}
}

func (a *AdminService) CreateSuperAdmin(ctx context.Context) error {
	// Смотрим есть ли уже созданный супер-админ в БД, если кол-во > 0, то не создаем нового супер-админа, если равно 0, то создаем супер-админа.
	// Проверяем длину слайса полученных админов, можно сделать отдельную функцию, которая будет возвращать кол-во админов с БД
	admins, err := a.repo.ListAdmins(ctx)
	if err != nil {
		a.log.Error("failed to get list of admins", zap.Error(err))
		return err
	}
	if len(admins) > 0 {
		a.log.Warn("super admin already exists")
		return nil
	}

	// Генерируем случайные логин и пароль для супер-админа
	creds := utils.GenerateAdminCredentials()

	// Создаём структуру для создания супер-админа
	admin := sqlc.CreateSuperAdminParams{
		UserName:     creds.Username,
		PasswordHash: creds.Password,
		CreatedBy:    sql.NullInt32{},
	}

	// Создаём новую запись с БД о новом админе
	_, err = a.repo.CreateSuperAdmin(ctx, admin)
	if err != nil {
		a.log.Error("failed to create admin", zap.Error(err))
		return err
	}

	a.log.Info("Super admin created successfully")

	utils.PrintCredentials(creds)
	return nil
}
