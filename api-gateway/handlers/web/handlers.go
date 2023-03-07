package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mksmstpck/to-rename/api-gateway/events"
	"github.com/nats-io/nats.go"
)

type Handlers struct {
	e          *echo.Echo
	conn       *nats.EncodedConn
	user       events.Users
	role       events.Roles
	permission events.Permissions
}

func NewHandlers(echo *echo.Echo, c *nats.EncodedConn, user *events.User, role *events.Role, permission *events.Permission) *Handlers {
	return &Handlers{
		e:          echo,
		conn:       c,
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
	user.GET("/id/:id", h.UserIdRead)
	user.GET("/username/:username", h.UserUsernameRead)
	user.GET("/email/:email", h.UserEmailRead)
	user.PUT("/", h.UserUpdate)
	user.DELETE("/:id", h.UserDelete)
	// role endpoints
	role.POST("/", h.RoleCreate)
	role.GET("/:id", h.RoleIdRead)
	role.PUT("/", h.RoleUpdate)
	role.DELETE("/:id", h.RoleDelete)
	// permission endpoints
	permission.POST("/", h.PermissionCreate)
	permission.GET("/:id", h.PermissionIdRead)
	permission.PUT("/", h.PermissionUpdate)
	permission.DELETE("/:id", h.PermissionDelete)
}
