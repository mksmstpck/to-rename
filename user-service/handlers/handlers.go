package handlers

import (
	"github.com/mkskstpck/to-rename/user-service/database"
	"github.com/nats-io/nats.go"
)

type Handler struct {
	conn *nats.EncodedConn
	user database.Users
}

func NewHandler(conn *nats.EncodedConn, user *database.UserDB) *Handler {
	return &Handler{
		conn: conn,
		user: user,
	}
}

func (h *Handler) HandleAll() {
	h.userIdRead()
	h.userUsernameRead()
	h.userEmailRead()
	h.userCreate()
	h.userUpdate()
	h.userDelete()
}
