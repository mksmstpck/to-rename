package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mksmstpck/to-rename/api-gateway/models"
)

func (w *Web) RoleCreate(c echo.Context) error {
	var r models.Role
	if err := c.Bind(&r); err != nil {
		return err
	}
	if err := c.Validate(&r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := w.r.RolePost(r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, r)
}

func (w *Web) RoleRead(c echo.Context) error {
	id := c.Param("id")
	r, err := w.r.RoleGet([]byte(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, r)
}

func (w *Web) RoleUpdate(c echo.Context) error {
	var r models.Role
	if err := c.Bind(&r); err != nil {
		return err
	}
	if err := c.Validate(&r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := w.r.RolePut(r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, r)
}

func (w *Web) RoleDelete(c echo.Context) error {
	id := c.Param("id")
	if err := w.r.RoleDelete([]byte(id)); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, id)
}
