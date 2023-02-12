package models

type Role struct {
	ID          string       `json:"id"`
	Name        string       `json:"name" validate:"required,name"`
	Permissions []Permission `json:"permissions" validate:"required,permissions"`
}
