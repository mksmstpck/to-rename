package models

type Role struct {
	ID            string   `json:"id"`
	Name          string   `json:"name" validate:"required,name"`
	PermissionIDs []string `json:"permissions" validate:"required,permissions"`
}
