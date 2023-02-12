package models

type Permission struct {
	ID        string   `json:"id"`
	Name      string   `json:"name" validate:"required,name"`
	AllowedTo []string `json:"allowedTo" validate:"required,allowedTo"`
}
