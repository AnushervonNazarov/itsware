package repositories

import (
	"context"
	"fmt"
	"itsware/internal/constants"
	"itsware/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDevice(ctx context.Context, device models.Device) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT devices.create(ROW(NULL::int, $1::citext, $2::citext, $3::int, $4::citext, NULL::timestamp, NULL::int, $5::int, $6::int, $7::int, $8::timestamp, $9::int, NULL::timestamp, NULL::int)::devices.device)`
	_, err := conn.Exec(ctx, query,
		device.Name,
		device.Description,
		device.DeviceStatusID,
		device.SerialNumber,
		device.CabinetID,
		device.TenantID,
		device.DeviceProfileID,
		device.CreatedOn,
		device.CreatedBy,
	)
	if err != nil {
		return fmt.Errorf("failed to create device: %w", err)
	}
	return nil
}

func GetDevice(ctx context.Context, id int) (*models.Device, error) {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return &models.Device{}, fmt.Errorf("database connection not found in context")
	}

	var device models.Device
	query := `SELECT * FROM devices.get_one($1)`
	row := conn.QueryRow(ctx, query, id)
	err := row.Scan(
		&device.ID, &device.Name, &device.Description, &device.DeviceStatusID, &device.SerialNumber,
		&device.CheckedOutOn, &device.CheckedOutBy, &device.CabinetID, &device.TenantID,
		&device.DeviceProfileID, &device.CreatedOn, &device.CreatedBy, &device.LastModifiedOn, &device.LastModifiedBy,
	)
	if err != nil {
		return &models.Device{}, err
	}
	return &device, nil
}

func GetAllDevices(ctx context.Context) ([]models.Device, error) {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return []models.Device{}, fmt.Errorf("database connection not found in context")
	}

	query := `SELECT * FROM devices.get_all()`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []models.Device
	for rows.Next() {
		var device models.Device
		if err := rows.Scan(
			&device.ID, &device.Name, &device.Description, &device.DeviceStatusID, &device.SerialNumber,
			&device.CheckedOutOn, &device.CheckedOutBy, &device.CabinetID, &device.TenantID,
			&device.DeviceProfileID, &device.CreatedOn, &device.CreatedBy,
			&device.LastModifiedOn, &device.LastModifiedBy); err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}

	return devices, nil
}

func UpdateDevice(ctx context.Context, device models.Device) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT devices.update(ROW($1::int, $2::citext, $3::citext, $4::int, $5::citext, NULL::timestamp, NULL::int, $6::int, NULL::int, $7::int, NULL::timestamp, NULL::int, NULL::timestamp, NULL::int)::devices.device)`
	_, err := conn.Exec(ctx, query,
		device.ID,
		device.Name,
		device.Description,
		device.DeviceStatusID,
		device.SerialNumber,
		device.CabinetID,
		device.DeviceProfileID,
	)
	if err != nil {
		return fmt.Errorf("failed to update device: %w", err)
	}
	return nil
}

func DeleteDevice(ctx context.Context, id int) error {
	conn, ok := ctx.Value(constants.DBConnKey).(*pgxpool.Conn)
	if !ok {
		return fmt.Errorf("database connection not found in context")
	}

	query := `SELECT devices.delete($1)`
	_, err := conn.Exec(ctx, query, id)
	return err
}
