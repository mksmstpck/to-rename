package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mksmstpck/to-rename/api-gateway/models"
)

func (h *Handlers) PermissionCreate(c echo.Context) error {
	p := new(models.Permission)
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Validate(&p); err != nil {
		return c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	if err := h.permission.PermissionPost(p); err != nil {
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusCreated, p)
}

func (h *Handlers) PermissionIdRead(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
	}
	p, err := h.permission.PermissionGet(int32(idInt))
	if p.ID == 0 {
		return c.JSON(http.StatusNotFound, models.Message{Message: "permission not found"})
	}
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, p)
}

func (h *Handlers) PermissionUpdate(c echo.Context) error {
	p := new(models.Permission)
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Validate(&p); err != nil {
		return c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	if err := h.permission.PermissionPut(p); err != nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusNoContent, nil)
}

func (h *Handlers) PermissionDelete(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
	}
	if err := h.permission.PermissionDelete(int32(idInt)); err != nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, models.Message{Message: id})
}
