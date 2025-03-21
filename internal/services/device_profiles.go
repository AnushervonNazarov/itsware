package services

import (
	"context"
	"itsware/internal/models"
	"itsware/internal/repositories"
)

func CreateDeviceProfile(ctx context.Context, deviceProfile models.DeviceProfile) error {
	return repositories.CreateDeviceProfile(ctx, deviceProfile)
}

func GetDeviceProfile(ctx context.Context, id int) (deviceProfile *models.DeviceProfile, err error) {
	return repositories.GetDeviceProfile(ctx, id)
}

func GetAllDeviceProfiles(ctx context.Context) (deviceProfile []models.DeviceProfile, err error) {
	return repositories.GetAllDeviceProfiles(ctx)
}

func UpdateDeviceProfile(ctx context.Context, deviceProfile models.DeviceProfile) error {
	return repositories.UpdateDeviceProfile(ctx, deviceProfile)
}

func DeleteDeviceProfile(ctx context.Context, id int) error {
	return repositories.DeleteDeviceProfile(ctx, id)
}
