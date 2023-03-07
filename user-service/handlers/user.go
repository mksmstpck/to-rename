package handlers

import (
	"github.com/mkskstpck/to-rename/user-service/models"
)

func (h *Handler) userIdRead() {
	h.conn.Subscribe("users-id-get", func(_, reply string, id int32) {
		user, err := h.user.UserFindOneId(id)
		if err != nil {
			res := models.Response{Status: "error", Message: err.Error()}
			h.conn.Publish(reply, res)
		}
		h.conn.Publish(reply, user)
	})
}

func (h *Handler) userUsernameRead() {
	h.conn.Subscribe("users-username-get", func(_, reply string, username string) {
		user, err := h.user.UserFindOneUsername(username)
		if err != nil {
			res := models.Response{Status: "error", Message: err.Error()}
			h.conn.Publish(reply, res)
		}
		h.conn.Publish(reply, user)
	})
}

func (h *Handler) userEmailRead() {
	h.conn.Subscribe("users-email-get", func(_, reply string, email string) {
		user, err := h.user.UserFindOneEmail(email)
		if err != nil {
			res := models.Response{Status: "error", Message: err.Error()}
			h.conn.Publish(reply, res)
		}
		h.conn.Publish(reply, user)
	})
}

func (h *Handler) userCreate() {
	h.conn.Subscribe("users-post", func(_, reply string, user models.User) {
		userExistUsername, err := h.user.UserFindOneUsername(user.Username)
		if err != nil {
			res := models.Response{Status: "error", Message: err.Error()}
			h.conn.Publish(reply, res)
		}
		if userExistUsername.ID != 0 {
			res := models.Response{Status: "error", Message: "user with this username already exists"}
			h.conn.Publish(reply, res)
		}
		userExistEmail, err := h.user.UserFindOneEmail(user.Email)
		if err != nil {
			res := models.Response{Status: "error", Message: err.Error()}
			h.conn.Publish(reply, res)
		}
		if userExistEmail.ID != 0 {
			res := models.Response{Status: "error", Message: "user with this email already exists"}
			h.conn.Publish(reply, res)
		}
		err = h.user.UserCreateOne(user)
		if err != nil {
			res := models.Response{Status: "error", Message: err.Error()}
			h.conn.Publish(reply, res)
		}
		h.conn.Publish(reply, models.Response{Status: "ok", Message: "created"})
	})
}

func (h *Handler) userUpdate() {
	h.conn.Subscribe("users-put", func(_, reply string, user models.User) {
		err := h.user.UserUpdateOne(user)
		if err != nil {
			res := models.Response{Status: "error", Message: err.Error()}
			h.conn.Publish(reply, res)
		}
		h.conn.Publish(reply, models.Response{Status: "ok", Message: "updated"})
	})
}

func (h *Handler) userDelete() {
	h.conn.Subscribe("users-delete", func(_, reply string, id int32) {
		err := h.user.UserDeleteOne(id)
		if err != nil {
			res := models.Response{Status: "error", Message: err.Error()}
			h.conn.Publish(reply, res)
		}
		h.conn.Publish(reply, models.Response{Status: "ok", Message: "deleted"})
	})
}
