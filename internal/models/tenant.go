package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Tenant struct {
	ID             int                `json:"id" validate:"required, uuid4"`
	Name           string             `json:"name" validate:"required,min=2,max=30"`
	IsSupport      bool               `json:"is_support" validate:"required"`
	CreatedOn      time.Time          `json:"created_on"`
	LastModifiedOn pgtype.Timestamptz `json:"last_modified_on"`
}
