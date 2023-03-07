package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mksmstpck/to-rename/api-gateway/models"
)

func (h *Handlers) RoleCreate(c echo.Context) error {
	r := new(models.Role)
	if err := c.Bind(&r); err != nil {
		return err
	}
	if err := c.Validate(&r); err != nil {
		return c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	if err := h.role.RolePost(r); err != nil {
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusCreated, r)
}

func (h *Handlers) RoleIdRead(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
	}
	r, err := h.role.RoleGet(int32(idInt))
	if r.ID == 0 {
		return c.JSON(http.StatusNotFound, models.Message{Message: "user not found"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, r)
}

func (h *Handlers) RoleUpdate(c echo.Context) error {
	r := new(models.Role)
	if err := c.Bind(&r); err != nil {
		return err
	}
	if err := c.Validate(&r); err != nil {
		return c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	if err := h.role.RolePut(r); err != nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusNoContent, nil)
}

func (h Handlers) RoleDelete(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
	}
	if err := h.role.RoleDelete(int32(idInt)); err != nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusNoContent, nil)
}
