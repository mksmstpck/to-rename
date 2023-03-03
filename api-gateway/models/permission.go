package models

type Permission struct {
	ID        int      `json:"id"`
	Name      string   `json:"name" validate:"required"`
	AllowedTo []string `json:"allowedTo" validate:"required"`
}
