package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mksmstpck/to-rename/api-gateway/models"
)

func (w *Web) PermissionCreate(c echo.Context) error {
	var p models.Permission
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Validate(&p); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := w.permission.PermissionPost(p); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, p)
}

func (w *Web) PermissionRead(c echo.Context) error {
	id := c.Param("id")
	p, err := w.permission.PermissionGet([]byte(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, p)
}

func (w *Web) PermissionUpdate(c echo.Context) error {
	var p models.Permission
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Validate(&p); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := w.permission.PermissionPut(p); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, p)
}

func (w *Web) PermissionDelete(c echo.Context) error {
	id := c.Param("id")
	if err := w.permission.PermissionDelete([]byte(id)); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, id)
}
