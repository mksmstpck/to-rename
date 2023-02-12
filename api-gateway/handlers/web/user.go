package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mksmstpck/to-rename/api-gateway/models"
)

func (w *Web) UserCreate(c echo.Context) error {
	var u models.User
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := w.user.UserPost(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, u)
}

func (w *Web) UserRead(c echo.Context) error {
	id := c.Param("id")
	u, err := w.user.UserGet([]byte(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, u)
}

func (w *Web) UserUpdate(c echo.Context) error {
	var u models.User
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := w.user.UserPut(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, u)
}

func (w *Web) UserDelete(c echo.Context) error {
	id := c.Param("id")
	if err := w.user.UserDelete([]byte(id)); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, id)
}
