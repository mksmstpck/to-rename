package events

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/mksmstpck/to-rename/api-gateway/models"
)

type RolePublisher interface {
	RoleGet(id []byte) (models.Role, error)
	RolePost(role models.Role) error
	RolePut(role models.Role) error
	RoleDelete(id []byte) error
}

func (p Pub) RoleGet(id []byte) (models.Role, error) {
	var role models.Role
	m, err := p.conn.Request("roles-get", id, 10*time.Millisecond)
	if err != nil {
		return role, err
	}
	sonic.Unmarshal(m.Data, &role)
	return role, nil
}

func (p Pub) RolePost(role models.Role) error {
	var res models.Response
	roleBytes, err := sonic.Marshal(role)
	if err != nil {
		return err
	}
	m, err := p.conn.Request("roles-post", roleBytes, 10*time.Millisecond)
	if err != nil {
		return err
	}
	sonic.Unmarshal(m.Data, &res)
	if res.Status == "ok" {
		return nil
	}
	return err
}

func (p Pub) RolePut(role models.Role) error {
	var res models.Response
	roleBytes, err := sonic.Marshal(role)
	if err != nil {
		return err
	}
	m, err := p.conn.Request("roles-put", roleBytes, 10*time.Millisecond)
	if err != nil {
		return err
	}
	sonic.Unmarshal(m.Data, &res)
	if res.Status == "ok" {
		return nil
	}
	return err
}

func (p Pub) RoleDelete(id []byte) error {
	var res models.Response
	m, err := p.conn.Request("roles-delete", id, 10*time.Millisecond)
	if err != nil {
		return err
	}
	sonic.Unmarshal(m.Data, &res)
	if res.Status == "ok" {
		return nil
	}
	return err
}
