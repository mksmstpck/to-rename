package handlers

import "github.com/labstack/echo/v4"

type Handler struct {
	e *echo.Echo
}

func NewHandler(echo *echo.Echo) *Handler {
	return &Handler{
		e: echo,
	}
}

func (h *Handler) hello(c echo.Context) error {
	return c.JSON(200, "Hello, World!")
}

func (h *Handler) All(e *echo.Echo) {
	h.e.GET("/", h.hello)
}
