package repositories

import (
	"context"
	"fmt"
	"itsware/internal/constants"
	"itsware/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTeam(ctx context.Context, team models.Team) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT teams.create(ROw(NULL::int, $1::citext, $2::int, $3::timestamp, $4::int, NULL::timestamp, NULL::int)::teams.team)`
	_, err := conn.Exec(ctx, query,
		team.Name,
		team.TenantID,
		team.CreatedOn,
		team.CreatedBy,
	)
	if err != nil {
		return fmt.Errorf("failed to create team: %w", err)
	}
	return nil
}

func GetTeam(ctx context.Context, id int) (*models.Team, error) {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return &models.Team{}, fmt.Errorf("database connection not found in context")
	}

	var team models.Team
	query := `SELECT * FROM teams.get_one($1)`
	row := conn.QueryRow(ctx, query, id)
	err := row.Scan(
		&team.ID, &team.Name, &team.TenantID, &team.CreatedOn, &team.CreatedBy,
		&team.LastModifiedOn, &team.LastModifiedBy,
	)
	if err != nil {
		return &models.Team{}, err
	}
	return &team, err
}

func GetAllTeams(ctx context.Context) ([]models.Team, error) {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return []models.Team{}, fmt.Errorf("database connection not found in context")
	}

	query := `SELECT * FROM teams.get_all()`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		if err := rows.Scan(
			&team.ID, &team.Name, &team.TenantID, &team.CreatedOn, &team.CreatedBy,
			&team.LastModifiedOn, &team.LastModifiedBy); err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}

	return teams, nil
}

func UpdateTeam(ctx context.Context, team models.Team) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT teams.update(ROW($1::int, $2::citext, NULL::int, NULL::timestamp, NULL::int, NULL::timestamp, NULL::int)::teams.team)`
	_, err := conn.Exec(ctx, query,
		team.ID,
		team.Name,
	)
	if err != nil {
		return fmt.Errorf("failed to update team: %w", err)
	}
	return nil
}

func DeleteTeam(ctx context.Context, id int) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT teams.delete($1)`
	_, err := conn.Exec(ctx, query, id)
	return err
}
