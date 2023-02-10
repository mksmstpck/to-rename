package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
)

type Web struct {
	e     *echo.Echo
	nconn *nats.Conn
}

func NewWeb(echo *echo.Echo, nc *nats.Conn) *Web {
	return &Web{
		e:     echo,
		nconn: nc,
	}
}

func (w *Web) All(e *echo.Echo) {
	u := w.e.Group("/users")
	u.POST("/", w.UserCreate)
}
