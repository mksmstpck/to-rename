package events

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/mksmstpck/to-rename/api-gateway/models"
)

type UserPublisher interface {
	UserGet(id []byte) (models.User, error)
	UserPost(user models.User) error
	UserPut(user models.User) error
	UserDelete(id []byte) error
}

func (p Pub) UserGet(id []byte) (models.User, error) {
	var user models.User
	m, err := p.conn.Request("users-get", id, 10*time.Millisecond)
	if err != nil {
		return user, err
	}
	sonic.Unmarshal(m.Data, &user)
	return user, nil
}

func (p Pub) UserPost(user models.User) error {
	var res models.Response
	userBytes, err := sonic.Marshal(user)
	if err != nil {
		return err
	}
	m, err := p.conn.Request("users-post", userBytes, 10*time.Millisecond)
	if err != nil {
		return err
	}
	sonic.Unmarshal(m.Data, &res)
	if res.Status == "ok" {
		return nil
	}
	return err
}

func (p Pub) UserPut(user models.User) error {
	var res models.Response
	userBytes, err := sonic.Marshal(user)
	if err != nil {
		return err
	}
	m, err := p.conn.Request("users-put", userBytes, 10*time.Millisecond)
	if err != nil {
		return err
	}
	sonic.Unmarshal(m.Data, &res)
	if res.Status == "ok" {
		return nil
	}
	return err
}

func (p Pub) UserDelete(id []byte) error {
	var res models.Response
	m, err := p.conn.Request("users-delete", id, 10*time.Millisecond)
	if err != nil {
		return err
	}
	sonic.Unmarshal(m.Data, &res)
	if res.Status == "ok" {
		return nil
	}
	return err
}
