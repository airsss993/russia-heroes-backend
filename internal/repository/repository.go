package repository

import (
	"context"

	"github.com/airsss993/russia-heroes-backend/internal/repository/postgres/sqlc"
)

// AdminRepo отвечает за работу с администраторами в базе данных
type AdminRepo interface {
	CreateAdmin(ctx context.Context, arg sqlc.CreateAdminParams) (sqlc.Admin, error)
	CreateSuperAdmin(ctx context.Context, arg sqlc.CreateSuperAdminParams) (sqlc.Admin, error)
	DeleteAdmin(ctx context.Context, id int32) error
	GetAdminByID(ctx context.Context, id int32) (sqlc.Admin, error)
	GetAdminByUsername(ctx context.Context, username string) (sqlc.Admin, error)
	ListAdmins(ctx context.Context) ([]sqlc.Admin, error)
}
