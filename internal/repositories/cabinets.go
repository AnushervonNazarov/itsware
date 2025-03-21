package repositories

import (
	"context"
	"fmt"
	"itsware/internal/constants"
	"itsware/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateCabinet(ctx context.Context, cabinet models.Cabinet) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT cabinets.create(ROW(NULL::int, $1::citext, $2::citext, $3::citext, $4::int, $5::timestamp, $6::int, NULL::timestamp, NULL::int)::cabinets.cabinet)`
	_, err := conn.Exec(ctx, query,
		cabinet.Name,
		cabinet.Location,
		cabinet.Description,
		cabinet.TenantID,
		cabinet.CreatedOn,
		cabinet.CreatedBy,
	)
	if err != nil {
		return fmt.Errorf("failed to create cabinet: %w", err)
	}
	return nil
}

func GetCabinet(ctx context.Context, id int) (*models.Cabinet, error) {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return &models.Cabinet{}, fmt.Errorf("database connection not found in context")
	}

	var cabinet models.Cabinet
	query := `SELECT * FROM cabinets.get_one($1)`
	row := conn.QueryRow(ctx, query, id)
	err := row.Scan(
		&cabinet.ID, &cabinet.Name, &cabinet.Location, &cabinet.Description,
		&cabinet.TenantID, &cabinet.CreatedOn, &cabinet.CreatedBy,
		&cabinet.LastModifiedOn, &cabinet.LastModifiedBy,
	)
	if err != nil {
		return &models.Cabinet{}, err
	}
	return &cabinet, err
}

func GetAllCabinets(ctx context.Context) ([]models.Cabinet, error) {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return []models.Cabinet{}, fmt.Errorf("database connection not found in context")
	}

	query := `SELECT * FROM cabinets.get_all()`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cabinets []models.Cabinet
	for rows.Next() {
		var cabinet models.Cabinet
		if err := rows.Scan(
			&cabinet.ID, &cabinet.Name, &cabinet.Location, &cabinet.Description,
			&cabinet.TenantID, &cabinet.CreatedOn, &cabinet.CreatedBy,
			&cabinet.LastModifiedOn, &cabinet.LastModifiedBy); err != nil {
			return nil, err
		}
		cabinets = append(cabinets, cabinet)
	}

	return cabinets, nil
}

func UpdateCabinet(ctx context.Context, cabinet models.Cabinet) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT cabinets.update(ROW($1::int, $2::citext, $3::citext, $4::citext, NULL::int, NULL::timestamp, NULL::int, NULL::timestamp, NULL::int)::cabinets.cabinet)`
	_, err := conn.Exec(ctx, query,
		cabinet.ID,
		cabinet.Name,
		cabinet.Location,
		cabinet.Description,
	)
	if err != nil {
		return fmt.Errorf("failed to update cabinet %w", err)
	}
	return nil
}

func DeleteCabinet(ctx context.Context, id int) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT cabinets.delete($1)`
	_, err := conn.Exec(ctx, query, id)
	return err
}
