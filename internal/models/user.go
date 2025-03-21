package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID             int                `json:"id" validate:"required,uuid4"`
	FirstName      string             `json:"first_name" validate:"required,min=2,max=50"`
	LastName       string             `json:"last_name" validate:"required,min=2,max=50"`
	Email          string             `json:"email" validate:"required,email"`
	Phone          string             `json:"phone" validate:"required,e164"`
	Password       string             `json:"password" validate:"required"`
	RoleID         int                `json:"role_id" validate:"required"`
	TenantID       int                `json:"tenant_id" validate:"required"`
	CreatedOn      time.Time          `json:"created_on"`
	CreatedBy      pgtype.Int4        `json:"created_by"`
	LastModifiedOn pgtype.Timestamptz `json:"last_modified_on"`
	LastModifiedBy pgtype.Int4        `json:"last_modified_by"`
}
