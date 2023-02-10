package models

type Permission struct {
	ID        string   `json:"id"`
	Name      string   `json:"name" validate:"required,name"`
	AllowedTo []string `json:"allowedTo" validate:"required,allowedTo"`
}

type Role struct {
	ID          string       `json:"id"`
	Name        string       `json:"name" validate:"required,name"`
	Permissions []Permission `json:"permissions" validate:"required,permissions"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username" validate:"required,name"`
	FullName string `json:"fullName" validate:"required,fullname"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
	Role     Role   `json:"role" validate:"required,role"`
}
