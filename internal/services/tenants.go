package services

import (
	"context"
	"itsware/internal/models"
	"itsware/internal/repositories"
)

func CreateTenant(ctx context.Context, tenant models.Tenant) error {
	return repositories.CreateTenant(ctx, tenant)
}

func GetTenant(ctx context.Context, id int) (*models.Tenant, error) {
	return repositories.GetTenant(ctx, id)
}

func GetAllTenants(ctx context.Context) ([]models.Tenant, error) {
	return repositories.GetAllTenants(ctx)
}

func UpdateTenant(ctx context.Context, tenant models.UpdateTenant) error {
	return repositories.UpdateTenant(ctx, tenant)
}

func DeleteTenant(ctx context.Context, id int) error {
	return repositories.DeleteTenant(ctx, id)
}
