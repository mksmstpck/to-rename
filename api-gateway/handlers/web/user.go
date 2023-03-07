package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mksmstpck/to-rename/api-gateway/models"
)

func (h *Handlers) UserIdRead(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
	}
	u, err := h.user.UserIdGet(int32(idInt))
	if u.ID == 0 {
		return c.JSON(http.StatusNotFound, models.Message{Message: "user not found"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Message{Message: "internal server error"})
	}
	return c.JSON(http.StatusOK, u)
}

func (h *Handlers) UserUsernameRead(c echo.Context) error {
	username := c.Param("username")
	u, err := h.user.UserUsernameGet(username)
	if u.ID == 0 {
		return c.JSON(http.StatusNotFound, models.Message{Message: "user not found"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Message{Message: "internal server error"})
	}
	return c.JSON(http.StatusOK, u)
}

func (h *Handlers) UserEmailRead(c echo.Context) error {
	email := c.Param("email")
	u, err := h.user.UserEmailGet(email)
	if u.ID == 0 {
		return c.JSON(http.StatusNotFound, models.Message{Message: "user not found"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Message{Message: "internal server error"})
	}
	return c.JSON(http.StatusOK, u)
}

func (h *Handlers) UserCreate(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	if err := h.user.UserPost(u); err != nil {
		return c.JSON(http.StatusConflict, models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, u)
}

func (h *Handlers) UserUpdate(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	if err := h.user.UserPut(u); err != nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusNoContent, nil)
}

func (h *Handlers) UserDelete(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
	}
	if err := h.user.UserDelete(int32(idInt)); err != nil {
		return c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusNoContent, nil)
}
