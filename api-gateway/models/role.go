package models

type Role struct {
	ID           int    `json:"id"`
	Name         string `json:"name" validate:"required"`
	PermissionID string `json:"permissions" validate:"required"`
}
