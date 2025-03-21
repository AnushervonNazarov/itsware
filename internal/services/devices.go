package services

import (
	"context"
	"itsware/internal/models"
	"itsware/internal/repositories"
)

func CreateDevice(ctx context.Context, device models.Device) error {
	return repositories.CreateDevice(ctx, device)
}

func GetDevice(ctx context.Context, id int) (*models.Device, error) {
	return repositories.GetDevice(ctx, id)
}

func GetAllDevices(ctx context.Context) ([]models.Device, error) {
	return repositories.GetAllDevices(ctx)
}

func UpdateDevice(ctx context.Context, device models.Device) error {
	return repositories.UpdateDevice(ctx, device)
}

func DeleteDevice(ctx context.Context, id int) error {
	return repositories.DeleteDevice(ctx, id)
}
