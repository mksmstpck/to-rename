package models

type Role struct {
	ID            int      `json:"id"`
	Name          string   `json:"name" validate:"required,name"`
	PermissionIDs []string `json:"permissions" validate:"required,permissions"`
}
