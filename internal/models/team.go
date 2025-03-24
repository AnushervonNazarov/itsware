package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Team struct {
	ID             int                `json:"id,omitempty" validate:"omitempty,uuid4"`
	Name           string             `json:"name" validate:"required,min=2,max=50"`
	TenantID       int                `json:"tenant_id" validate:"required"`
	CreatedOn      time.Time          `json:"created_on"`
	CreatedBy      int                `json:"created_by"`
	LastModifiedOn pgtype.Timestamptz `json:"last_modified_on"`
	LastModifiedBy pgtype.Int4        `json:"last_modified_by"`
}

type UpdateTeam struct {
	ID   int    `json:"id,omitempty" validate:"omitempty"`
	Name string `json:"name" validate:"required,min=2,max=50"`
}
