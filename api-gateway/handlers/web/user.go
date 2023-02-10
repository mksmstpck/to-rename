package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mksmstpck/to-rename/api-gateway/events"
	"github.com/mksmstpck/to-rename/api-gateway/models"
)

func (w *Web) UserCreate(c echo.Context) error {
	u := models.User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := events.NewPub(w.nconn).WriteUser(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, u)
}

func (w *Web) UserGet(c echo.Context) error {
	id := c.Param("id")
	u, err := events.NewPub(w.nconn).GetUser([]byte(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, u)
}
