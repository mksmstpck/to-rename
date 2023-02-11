package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mksmstpck/to-rename/api-gateway/events"
	"github.com/nats-io/nats.go"
)

type Web struct {
	e     *echo.Echo
	nconn *nats.Conn
	u     events.UserPublisher
}

func NewWeb(echo *echo.Echo, nc *nats.Conn, pub *events.Pub) *Web {
	return &Web{
		e:     echo,
		nconn: nc,
		u:     pub,
	}
}

func (w *Web) All() {
	u := w.e.Group("/users")
	u.POST("/", w.UserCreate)
	u.GET("/:id", w.UserGet)
	u.PUT("/", w.UserUpdate)
	u.DELETE("/:id", w.UserDelete)
}
