package repository

import (
	"context"

	"github.com/airsss993/russia-heroes-backend/internal/repository/postgres/sqlc"
)

type adminRepo struct {
	queries *sqlc.Queries
}

func NewAdminRepo(queries *sqlc.Queries) AdminRepo {
	return &adminRepo{
		queries: queries,
	}
}

func (r *adminRepo) CreateAdmin(ctx context.Context, params sqlc.CreateAdminParams) (sqlc.Admin, error) {
	return r.queries.CreateAdmin(ctx, params)
}

func (r *adminRepo) CreateSuperAdmin(ctx context.Context, params sqlc.CreateSuperAdminParams) (sqlc.Admin, error) {
	return r.queries.CreateSuperAdmin(ctx, params)
}

func (r *adminRepo) GetAdminByID(ctx context.Context, id int32) (sqlc.Admin, error) {
	return r.queries.GetAdminByID(ctx, id)
}

func (r *adminRepo) GetAdminByUsername(ctx context.Context, username string) (sqlc.Admin, error) {
	return r.queries.GetAdminByUsername(ctx, username)
}

func (r *adminRepo) ListAdmins(ctx context.Context) ([]sqlc.Admin, error) {
	return r.queries.ListAdmins(ctx)
}

func (r *adminRepo) DeleteAdmin(ctx context.Context, id int32) error {
	return r.queries.DeleteAdmin(ctx, id)
}
