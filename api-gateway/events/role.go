package events

import (
	"errors"
	"time"

	"github.com/mksmstpck/to-rename/api-gateway/models"
)

func (r Role) RoleGet(id int32) (models.Role, error) {
	var role models.Role
	err := r.conn.Request("roles-get", id, &role, time.Second)
	if err != nil {
		return role, err
	}
	return role, nil
}

func (r Role) RolePost(role *models.Role) error {
	var res models.Response
	err := r.conn.Request("roles-post", role, &res, time.Second)
	if err != nil {
		return err
	}
	if res.Status == "ok" {
		return nil
	} else {
		return errors.New(res.Message)
	}
}

func (r Role) RolePut(role *models.Role) error {
	var res models.Response
	err := r.conn.Request("roles-put", role, &res, time.Second)
	if err != nil {
		return err
	}
	if res.Status == "ok" {
		return nil
	} else {
		return errors.New(res.Message)
	}
}

func (r Role) RoleDelete(id int32) error {
	var res models.Response
	err := r.conn.Request("roles-delete", id, &res, time.Second)
	if err != nil {
		return err
	}
	if res.Status == "ok" {
		return nil
	} else {
		return errors.New(res.Message)
	}
}
