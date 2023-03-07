package events

import (
	"errors"
	"time"

	"github.com/mksmstpck/to-rename/api-gateway/models"
)

func (p Permission) PermissionGet(id int32) (models.Permission, error) {
	var permission models.Permission
	err := p.conn.Request("permission-get", id, &permission, 10*time.Millisecond)
	if err != nil {
		return permission, err
	}
	return permission, nil
}

func (p Permission) PermissionPost(permission *models.Permission) error {
	var res models.Response
	err := p.conn.Request("permission-post", permission, &res, 10*time.Millisecond)
	if err != nil {
		return err
	}
	if res.Status == "ok" {
		return nil
	} else {
		return errors.New(res.Message)
	}
}

func (p Permission) PermissionPut(permission *models.Permission) error {
	var res models.Response
	err := p.conn.Request("permission-put", permission, &res, 10*time.Millisecond)
	if err != nil {
		return err
	}
	if res.Status == "ok" {
		return nil
	} else {
		return errors.New(res.Message)
	}
}

func (p Permission) PermissionDelete(id int32) error {
	var res models.Response
	err := p.conn.Request("permission-delete", id, &res, 10*time.Millisecond)
	if err != nil {
		return err
	}
	if res.Status == "ok" {
		return nil
	} else {
		return errors.New(res.Message)
	}
}
