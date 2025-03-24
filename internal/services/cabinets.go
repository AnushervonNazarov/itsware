package services

import (
	"context"
	"itsware/internal/models"
	"itsware/internal/repositories"
)

func CreateCabinet(ctx context.Context, cabinet models.Cabinet) error {
	return repositories.CreateCabinet(ctx, cabinet)
}

func GetCabinet(ctx context.Context, id int) (cabinet *models.Cabinet, err error) {
	return repositories.GetCabinet(ctx, id)
}

func GetAllCabinets(ctx context.Context) (cabinet []models.Cabinet, err error) {
	return repositories.GetAllCabinets(ctx)
}

func UpdateCabinet(ctx context.Context, cabinet models.UpdateCabinet) error {
	return repositories.UpdateCabinet(ctx, cabinet)
}

func DeleteCabinet(ctx context.Context, id int) error {
	return repositories.DeleteCabinet(ctx, id)
}

func AddCabinetToTeam(ctx context.Context, teamCabinet models.TeamCabinet) error {
	return repositories.AddCabinetToTeam(ctx, teamCabinet)
}

func RemoveCabinetFromTeam(ctx context.Context, cabinet_id int, team_id int) error {
	return repositories.RemoveCabinetFromTeam(ctx, cabinet_id, team_id)
}
