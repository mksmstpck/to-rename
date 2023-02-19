package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mksmstpck/to-rename/api-gateway/events"
	"github.com/nats-io/nats.go"
)

type Handlers struct {
	e          *echo.Echo
	nconn      *nats.Conn
	user       events.Users
	role       events.Roles
	permission events.Permissions
}

func NewHandlers(echo *echo.Echo, nc *nats.Conn, user *events.User, role *events.Role, permission *events.Permission) *Handlers {
	return &Handlers{
		e:          echo,
		nconn:      nc,
		user:       user,
		role:       role,
		permission: permission,
	}
}

func (h *Handlers) All() {
	// grougs
	user := h.e.Group("/users")
	role := h.e.Group("/roles")
	permission := h.e.Group("/permissions")
	// user endpoints
	user.POST("/", h.UserCreate)
	user.GET("/:id", h.UserRead)
	user.PUT("/", h.UserUpdate)
	user.DELETE("/:id", h.UserDelete)
	// role endpoints
	role.POST("/", h.RoleCreate)
	role.GET("/:id", h.RoleRead)
	role.PUT("/", h.RoleUpdate)
	role.DELETE("/:id", h.RoleDelete)
	// permission endpoints
	permission.POST("/", h.PermissionCreate)
	permission.GET("/:id", h.PermissionRead)
	permission.PUT("/", h.PermissionUpdate)
	permission.DELETE("/:id", h.PermissionDelete)
}
