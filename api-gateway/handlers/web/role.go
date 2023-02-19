package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mksmstpck/to-rename/api-gateway/models"
)

func (h *Handlers) RoleCreate(c echo.Context) error {
	var r models.Role
	if err := c.Bind(&r); err != nil {
		return err
	}
	if err := c.Validate(&r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.role.RolePost(r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, r)
}

func (h *Handlers) RoleRead(c echo.Context) error {
	id := c.Param("id")
	r, err := h.role.RoleGet([]byte(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, r)
}

func (h *Handlers) RoleUpdate(c echo.Context) error {
	var r models.Role
	if err := c.Bind(&r); err != nil {
		return err
	}
	if err := c.Validate(&r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.role.RolePut(r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, r)
}

func (h Handlers) RoleDelete(c echo.Context) error {
	id := c.Param("id")
	if err := h.role.RoleDelete([]byte(id)); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, id)
}
