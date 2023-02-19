package events

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/mksmstpck/to-rename/api-gateway/models"
)

func (p Permission) PermissionGet(id []byte) (models.Permission, error) {
	var permission models.Permission
	m, err := p.conn.Request("permission-get", id, 10*time.Millisecond)
	if err != nil {
		return permission, err
	}
	sonic.Unmarshal(m.Data, &permission)
	return permission, nil
}

func (p Permission) PermissionPost(permission models.Permission) error {
	var res models.Response
	permissionBytes, err := sonic.Marshal(permission)
	if err != nil {
		return err
	}
	m, err := p.conn.Request("permission-post", permissionBytes, 10*time.Millisecond)
	if err != nil {
		return err
	}
	sonic.Unmarshal(m.Data, &res)
	if res.Status == "ok" {
		return nil
	}
	return err
}

func (p Permission) PermissionPut(permission models.Permission) error {
	var res models.Response
	permissionBytes, err := sonic.Marshal(permission)
	if err != nil {
		return err
	}
	m, err := p.conn.Request("permission-put", permissionBytes, 10*time.Millisecond)
	if err != nil {
		return err
	}
	sonic.Unmarshal(m.Data, &res)
	if res.Status == "ok" {
		return nil
	}
	return err
}

func (p Permission) PermissionDelete(id []byte) error {
	var res models.Response
	m, err := p.conn.Request("permission-delete", id, 10*time.Millisecond)
	if err != nil {
		return err
	}
	sonic.Unmarshal(m.Data, &res)
	if res.Status == "ok" {
		return nil
	}
	return err
}
