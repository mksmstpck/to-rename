package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mksmstpck/to-rename/api-gateway/events"
	"github.com/nats-io/nats.go"
)

type Web struct {
	e          *echo.Echo
	nconn      *nats.Conn
	user       events.UserPublisher
	role       events.RolePublisher
	permission events.PermissionPublisher
}

func NewWeb(echo *echo.Echo, nc *nats.Conn, pub *events.Pub) *Web {
	return &Web{
		e:          echo,
		nconn:      nc,
		user:       pub,
		role:       pub,
		permission: pub,
	}
}

func (w *Web) All() {
	user := w.e.Group("/users")
	role := w.e.Group("/roles")
	permission := w.e.Group("/permissions")
	// user endpoints
	user.POST("/", w.UserCreate)
	user.GET("/:id", w.UserRead)
	user.PUT("/", w.UserUpdate)
	user.DELETE("/:id", w.UserDelete)
	// role endpoints
	role.POST("/", w.RoleCreate)
	role.GET("/:id", w.RoleRead)
	role.PUT("/", w.RoleUpdate)
	role.DELETE("/:id", w.RoleDelete)
	// permission endpoints
	permission.POST("/", w.PermissionCreate)
	permission.GET("/:id", w.PermissionRead)
	permission.PUT("/", w.PermissionUpdate)
	permission.DELETE("/:id", w.PermissionDelete)
}
