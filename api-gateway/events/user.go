package events

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/mksmstpck/to-rename/api-gateway/models"
)

func (p *Pub) GetUser(id []byte) (models.User, error) {
	var user models.User
	m, err := p.conn.Request("users", id, 10*time.Millisecond)
	if err != nil {
		panic(err)
	}
	sonic.Unmarshal(m.Data, &user)
	return user, nil
}

func (p *Pub) WriteUser(user models.User) error {
	var res models.Response
	userBytes, err := sonic.Marshal(user)
	if err != nil {
		return err
	}
	m, err := p.conn.Request("users", userBytes, 10*time.Millisecond)
	if err != nil {
		return err
	}
	sonic.Unmarshal(m.Data, &res)
	if res.Status == "ok" {
		return nil
	}
	if res.Status != "ok" {
		return err
	}
	return nil
}
