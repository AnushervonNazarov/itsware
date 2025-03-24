package services

import (
	"context"
	"itsware/internal/models"
	"itsware/internal/repositories"
)

func CreateTeam(ctx context.Context, team models.Team) error {
	return repositories.CreateTeam(ctx, team)
}

func GetTeam(ctx context.Context, id int) (team *models.Team, err error) {
	return repositories.GetTeam(ctx, id)
}

func GetAllTeams(ctx context.Context) (team []models.Team, err error) {
	return repositories.GetAllTeams(ctx)
}

func UpdateTeam(ctx context.Context, team models.UpdateTeam) error {
	return repositories.UpdateTeam(ctx, team)
}

func DeleteTeam(ctx context.Context, id int) error {
	return repositories.DeleteTeam(ctx, id)
}
