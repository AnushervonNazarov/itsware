package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type TeamCabinet struct {
	CabinetID      int                `json:"cabinet_id" validate:"required"`
	TeamID         int                `json:"team_id" validate:"required"`
	CreatedOn      time.Time          `json:"created_on"`
	CreatedBy      int                `json:"created_by"`
	LastModifiedOn pgtype.Timestamptz `json:"last_modified_on"`
	LastModifiedBy pgtype.Int4        `json:"last_modified_by"`
}
