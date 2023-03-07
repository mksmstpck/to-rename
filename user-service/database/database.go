package database

import (
	"database/sql"

	"github.com/mkskstpck/to-rename/user-service/models"
)

type UserDB struct {
	database *sql.DB
}

func NewUserDB(database *sql.DB) *UserDB {
	return &UserDB{
		database: database,
	}
}

type Users interface {
	UserFindOneId(ID int32) (models.User, error)
	UserFindOneUsername(username string) (models.User, error)
	UserFindOneEmail(email string) (models.User, error)
	UserCreateOne(user models.User) error
	UserUpdateOne(user models.User) error
	UserDeleteOne(ID int32) error
}

type Database struct {
	User Users
}
