package events

import (
	"errors"
	"time"

	"github.com/mksmstpck/to-rename/api-gateway/models"
)

func (u User) UserEmailGet(email string) (models.User, error) {
	var user models.User
	err := u.conn.Request("users-email-get", email, &user, time.Second)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u User) UserUsernameGet(username string) (models.User, error) {
	var user models.User
	err := u.conn.Request("users-username-get", username, &user, time.Second)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u User) UserIdGet(id int32) (models.User, error) {
	var user models.User
	err := u.conn.Request("users-id-get", id, &user, time.Second)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u User) UserPost(user *models.User) error {
	var res models.Response
	err := u.conn.Request("users-post", user, &res, time.Second)
	if err != nil {
		return err
	}
	if res.Status == "ok" {
		return nil
	} else {
		return errors.New(res.Message)
	}
}

func (u User) UserPut(user *models.User) error {
	var res models.Response
	err := u.conn.Request("users-put", user, &res, time.Second)
	if err != nil {
		return err
	}
	if res.Status == "ok" {
		return nil
	} else {
		return errors.New(res.Message)
	}
}

func (u User) UserDelete(id int32) error {
	var res models.Response
	err := u.conn.Request("users-delete", id, &res, time.Second)
	if err != nil {
		return err
	}
	if res.Status == "ok" {
		return nil
	} else {
		return errors.New(res.Message)
	}
}
