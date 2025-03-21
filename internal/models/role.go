package models

type Role struct {
	ID   int    `json:"id" validate:"required,uuid4"`
	Name string `json:"name" validate:"required,min=2,max=20"`
}
