package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mksmstpck/to-rename/api-gateway/models"
)

func (h *Handlers) PermissionCreate(c echo.Context) error {
	var p models.Permission
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Validate(&p); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.permission.PermissionPost(p); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, p)
}

func (h *Handlers) PermissionRead(c echo.Context) error {
	id := c.Param("id")
	p, err := h.permission.PermissionGet([]byte(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, p)
}

func (h *Handlers) PermissionUpdate(c echo.Context) error {
	var p models.Permission
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Validate(&p); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.permission.PermissionPut(p); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, p)
}

func (h *Handlers) PermissionDelete(c echo.Context) error {
	id := c.Param("id")
	if err := h.permission.PermissionDelete([]byte(id)); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, id)
}
