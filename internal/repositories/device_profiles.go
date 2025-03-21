package repositories

import (
	"context"
	"fmt"
	"itsware/internal/constants"
	"itsware/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDeviceProfile(ctx context.Context, deviceProfile models.DeviceProfile) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT device_profiles.create(ROW(NULL::int, $1::citext, $2::citext, $3::int, $4::timestamp, $5::int, NULL::timestamp, NULL::int)::device_profiles.device_profile)`
	_, err := conn.Exec(ctx, query,
		deviceProfile.Name,
		deviceProfile.Description,
		deviceProfile.TenantID,
		deviceProfile.CreatedOn,
		deviceProfile.CreatedBy,
	)
	if err != nil {
		return fmt.Errorf("failed to create device profile: %w", err)
	}
	return nil
}

func GetDeviceProfile(ctx context.Context, id int) (*models.DeviceProfile, error) {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return &models.DeviceProfile{}, fmt.Errorf("database connection not found in context")
	}

	var deviceProfile models.DeviceProfile
	query := `SELECT * FROM device_profiles.get_one($1)`
	row := conn.QueryRow(ctx, query, id)
	err := row.Scan(
		&deviceProfile.ID, &deviceProfile.Name, &deviceProfile.Description, &deviceProfile.TenantID,
		&deviceProfile.CreatedOn, &deviceProfile.CreatedBy, &deviceProfile.LastModifiedOn, &deviceProfile.LastModifiedBy,
	)
	if err != nil {
		return &models.DeviceProfile{}, err
	}
	return &deviceProfile, nil
}

func GetAllDeviceProfiles(ctx context.Context) ([]models.DeviceProfile, error) {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return []models.DeviceProfile{}, fmt.Errorf("database connection not found in context")
	}

	query := `SELECT * FROM device_profiles.get_all()`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deviceProfiles []models.DeviceProfile
	for rows.Next() {
		var deviceProfile models.DeviceProfile
		if err := rows.Scan(
			&deviceProfile.ID, &deviceProfile.Name, &deviceProfile.Description, &deviceProfile.TenantID,
			&deviceProfile.CreatedOn, &deviceProfile.CreatedBy,
			&deviceProfile.LastModifiedOn, &deviceProfile.LastModifiedBy); err != nil {
			return nil, err
		}
		deviceProfiles = append(deviceProfiles, deviceProfile)
	}

	return deviceProfiles, nil
}

func UpdateDeviceProfile(ctx context.Context, deviceProfile models.DeviceProfile) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT device_profiles.update(ROW($1::int, $2::citext, $3::citext, NULL::int, NULL::timestamp, NULL::int, NULL::timestamp, NULL::int)::device_profiles.device_profile)`
	_, err := conn.Exec(ctx, query,
		deviceProfile.ID,
		deviceProfile.Name,
		deviceProfile.Description,
	)
	if err != nil {
		return fmt.Errorf("failed to update device profile: %w", err)
	}
	return nil
}

func DeleteDeviceProfile(ctx context.Context, id int) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT device_profiles.delete($1)`
	_, err := conn.Exec(ctx, query, id)
	return err
}
