package events

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/mksmstpck/to-rename/api-gateway/models"
)

func (r Role) RoleGet(id []byte) (models.Role, error) {
	var role models.Role
	m, err := r.conn.Request("roles-get", id, 10*time.Millisecond)
	if err != nil {
		return role, err
	}
	sonic.Unmarshal(m.Data, &role)
	return role, nil
}

func (r Role) RolePost(role models.Role) error {
	var res models.Response
	roleBytes, err := sonic.Marshal(role)
	if err != nil {
		return err
	}
	m, err := r.conn.Request("roles-post", roleBytes, 10*time.Millisecond)
	if err != nil {
		return err
	}
	sonic.Unmarshal(m.Data, &res)
	if res.Status == "ok" {
		return nil
	}
	return err
}

func (r Role) RolePut(role models.Role) error {
	var res models.Response
	roleBytes, err := sonic.Marshal(role)
	if err != nil {
		return err
	}
	m, err := r.conn.Request("roles-put", roleBytes, 10*time.Millisecond)
	if err != nil {
		return err
	}
	sonic.Unmarshal(m.Data, &res)
	if res.Status == "ok" {
		return nil
	}
	return err
}

func (r Role) RoleDelete(id []byte) error {
	var res models.Response
	m, err := r.conn.Request("roles-delete", id, 10*time.Millisecond)
	if err != nil {
		return err
	}
	sonic.Unmarshal(m.Data, &res)
	if res.Status == "ok" {
		return nil
	}
	return err
}
