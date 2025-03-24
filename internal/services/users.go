package services

import (
	"context"
	"itsware/internal/models"
	"itsware/internal/repositories"
)

type User struct {
	Repo *repositories.User
}

func CreateUser(ctx context.Context, user models.User) error {
	return repositories.CreateUser(ctx, user)
}

func GetUser(ctx context.Context, id int) (user *models.User, err error) {
	return repositories.GetUser(ctx, id)
}

func GetAllUsers(ctx context.Context) (user []models.User, err error) {
	return repositories.GetAllUsers(ctx)
}

func UpdateUser(ctx context.Context, user models.UpdateUser) error {
	return repositories.UpdateUser(ctx, user)
}

func DeleteUser(ctx context.Context, id int) error {
	return repositories.DeleteUser(ctx, id)
}

func AddUserToTeam(ctx context.Context, teamUser models.TeamUser) error {
	return repositories.AddUserToTeam(ctx, teamUser)
}

func RemoveUserFromTeam(ctx context.Context, user_id int, team_id int) error {
	return repositories.RemoveUserFromTeam(ctx, user_id, team_id)
}
