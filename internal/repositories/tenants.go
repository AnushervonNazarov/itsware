package repositories

import (
	"context"
	"fmt"
	"itsware/internal/constants"
	"itsware/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTenant(ctx context.Context, tenant models.Tenant) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT tenants.create(ROW(NULL::int, $1::citext, $2::boolean, $3::timestamp, NULL::timestamp)::tenants.tenant)`
	_, err := conn.Exec(ctx, query,
		tenant.Name,
		tenant.IsSupport,
		tenant.CreatedOn,
	)
	if err != nil {
		return fmt.Errorf("failed to create tenant: %w", err)
	}
	return nil
}

func GetTenant(ctx context.Context, id int) (*models.Tenant, error) {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return &models.Tenant{}, fmt.Errorf("database connection not found in context")
	}

	var tenant models.Tenant

	query := `SELECT * FROM tenants.get_one($1)`
	row := conn.QueryRow(ctx, query, id)
	err := row.Scan(
		&tenant.ID, &tenant.Name, &tenant.IsSupport, &tenant.CreatedOn, &tenant.LastModifiedOn,
	)
	if err != nil {
		return &models.Tenant{}, err
	}
	return &tenant, nil
}

func GetAllTenants(ctx context.Context) ([]models.Tenant, error) {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return []models.Tenant{}, fmt.Errorf("database connection not found in context")
	}

	query := `SELECT * FROM tenants.get_all()`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenants []models.Tenant
	for rows.Next() {
		var tenant models.Tenant
		if err := rows.Scan(
			&tenant.ID, &tenant.Name, &tenant.IsSupport, &tenant.CreatedOn,
			&tenant.LastModifiedOn); err != nil {
			return nil, err
		}
		tenants = append(tenants, tenant)
	}

	return tenants, nil
}

func UpdateTenant(ctx context.Context, tenant models.Tenant) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT tenants.update(ROW($1::int, $2::citext, $3::boolean, NULL::timestamp, NULL::timestamp)::tenants.tenant)`
	_, err := conn.Exec(ctx, query,
		tenant.ID,
		tenant.Name,
		tenant.IsSupport,
	)
	if err != nil {
		return fmt.Errorf("failed to update tenant: %w", err)
	}
	return nil
}

func DeleteTenant(ctx context.Context, id int) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT tenants.delete($1)`
	_, err := conn.Exec(ctx, query, id)
	return err
}
