package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Cabinet struct {
	ID             int                `json:"id,omitempty" validate:"omitempty,uuid4"`
	Name           string             `json:"name" validate:"required,min=2,max=50"`
	Location       string             `json:"location" validate:"required"`
	Description    string             `json:"description" validate:"required"`
	TenantID       int                `json:"tenant_id" validate:"required"`
	CreatedOn      time.Time          `json:"created_on"`
	CreatedBy      pgtype.Int4        `json:"created_by"`
	LastModifiedOn pgtype.Timestamptz `json:"last_modified_on"`
	LastModifiedBy pgtype.Int4        `json:"last_modified_by"`
}

type UpdateCabinet struct {
	ID          int    `json:"id,omitempty" validate:"omitempty"`
	Name        string `json:"name" validate:"required,min=2,max=50"`
	Location    string `json:"location" validate:"required"`
	Description string `json:"description" validate:"required"`
}
