package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Device struct {
	ID              int                `json:"id,omitempty" validate:"omitempty,uuid4"`
	Name            string             `json:"name" validate:"required,min=2,max=50"`
	Description     string             `json:"description" validate:"required"`
	DeviceStatusID  int                `json:"device_status_id" validate:"required"`
	SerialNumber    string             `json:"serial_number" validate:"required"`
	CheckedOutOn    pgtype.Timestamptz `json:"checked_out_on"`
	CheckedOutBy    pgtype.Int4        `json:"checked_out_by"`
	CabinetID       int                `json:"cabinet_id" validate:"required"`
	TenantID        int                `json:"tenant_id" validate:"required"`
	DeviceProfileID int                `json:"device_profile_id" validate:"required"`
	CreatedOn       time.Time          `json:"created_on"`
	CreatedBy       pgtype.Int4        `json:"created_by"`
	LastModifiedOn  pgtype.Timestamptz `json:"last_modified_on"`
	LastModifiedBy  pgtype.Int4        `json:"last_modified_by"`
}

type UpdateDevice struct {
	ID              int    `json:"id,omitempty" validate:"omitempty"`
	Name            string `json:"name" validate:"required,min=2,max=50"`
	Description     string `json:"description" validate:"required"`
	DeviceStatusID  int    `json:"device_status_id" validate:"required"`
	SerialNumber    string `json:"serial_number" validate:"required"`
	CabinetID       int    `json:"cabinet_id" validate:"required"`
	DeviceProfileID int    `json:"device_profile_id" validate:"required"`
}
