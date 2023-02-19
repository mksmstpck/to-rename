package events

import (
	"github.com/mksmstpck/to-rename/api-gateway/models"
	"github.com/nats-io/nats.go"
)

type User struct {
	conn *nats.Conn
}

func NewUserEvent(conn *nats.Conn) *User {
	return &User{
		conn: conn,
	}
}

type Permission struct {
	conn *nats.Conn
}

func NewPermissionEvent(conn *nats.Conn) *Permission {
	return &Permission{
		conn: conn,
	}
}

type Role struct {
	conn *nats.Conn
}

func NewRoleEvent(conn *nats.Conn) *Role {
	return &Role{
		conn: conn,
	}
}

type Users interface {
	UserGet(id []byte) (models.User, error)
	UserPost(user models.User) error
	UserPut(user models.User) error
	UserDelete(id []byte) error
}

type Permissions interface {
	PermissionGet(id []byte) (models.Permission, error)
	PermissionPost(permission models.Permission) error
	PermissionPut(permission models.Permission) error
	PermissionDelete(id []byte) error
}

type Roles interface {
	RoleGet(id []byte) (models.Role, error)
	RolePost(role models.Role) error
	RolePut(role models.Role) error
	RoleDelete(id []byte) error
}

type Events struct {
	User       Users
	Permission Permissions
	Role       Roles
}
