package events

import (
	"log"
	"time"

	"github.com/bytedance/sonic"
	"github.com/mksmstpck/to-rename/api-gateway/models"
)

func (u User) UserGet(id []byte) (models.User, error) {
	var user models.User
	log.Println("user send request")
	m, err := u.conn.Request("users-get", id, 60*time.Second)
	if err != nil {
		return user, err
	}
	sonic.Unmarshal(m.Data, &user)
	return user, nil
}

func (u User) UserPost(user models.User) error {
	//	var res models.Response
	userBytes, err := sonic.Marshal(&user)
	if err != nil {
		return err
	}
	_, err = u.conn.Request("users-post", userBytes, 10*time.Millisecond)
	if err != nil {
		return err
	}
	//	sonic.Unmarshal(m.Data, &res)
	//	if res.Status == "ok" {
	//		return nil
	//	}
	return err
}

func (u User) UserPut(user models.User) error {
	var res models.Response
	userBytes, err := sonic.Marshal(user)
	if err != nil {
		return err
	}
	m, err := u.conn.Request("users-put", userBytes, 10*time.Millisecond)
	if err != nil {
		return err
	}
	sonic.Unmarshal(m.Data, &res)
	if res.Status == "ok" {
		return nil
	}
	return err
}

func (u User) UserDelete(id []byte) error {
	var res models.Response
	m, err := u.conn.Request("users-delete", id, 10*time.Millisecond)
	if err != nil {
		return err
	}
	sonic.Unmarshal(m.Data, &res)
	if res.Status == "ok" {
		return nil
	}
	return err
}
