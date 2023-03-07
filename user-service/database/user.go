package database

import (
	"database/sql"
	"errors"

	"github.com/mkskstpck/to-rename/user-service/models"
	"github.com/mkskstpck/to-rename/user-service/services"
)

func (d *UserDB) UserFindOneId(ID int32) (models.User, error) {
	var user models.User
	selectQuery := `select "id", "username", "fullname", "email", "roleid" from "Users" where "id" = $1`
	row := d.database.QueryRow(selectQuery, ID)
	err := row.Scan(&user.ID, &user.Username, &user.FullName, &user.Email, &user.RoleID)
	if err != sql.ErrNoRows {
		return user, err
	}
	return user, nil
}

func (d *UserDB) UserFindOneEmail(email string) (models.User, error) {
	var user models.User
	selectQuery := `select "id", "username", "fullname", "email", "roleid" from "Users" where "email" = $1`
	row := d.database.QueryRow(selectQuery, email)
	err := row.Scan(&user.ID, &user.Username, &user.FullName, &user.Email, &user.RoleID)
	if err != sql.ErrNoRows {
		return user, err
	}
	return user, nil
}

func (d *UserDB) UserFindOneUsername(username string) (models.User, error) {
	var user models.User
	selectQuery := `select "id", "username", "fullname", "email", "roleid" from "Users" where "username" = $1`
	row := d.database.QueryRow(selectQuery, username)
	err := row.Scan(&user.ID, &user.Username, &user.FullName, &user.Email, &user.RoleID)
	if err != sql.ErrNoRows {
		return user, err
	}
	return user, nil
}

func (d *UserDB) UserCreateOne(user models.User) error {
	insertQuery := `insert into "Users"("username", "fullname", "email", "password", "roleid") values ($1, $2, $3, $4, $5)`
	_, err := d.database.Exec(
		insertQuery,
		user.Username,
		user.FullName,
		user.Email,
		services.PasswordHash(user.Password),
		user.RoleID)
	if err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (d *UserDB) UserUpdateOne(user models.User) error {
	u, err := d.UserFindOneId(user.ID)
	if u.ID == 0 {
		return errors.New("user not found")
	}
	updateQuery := `update "Users" set "fullname" = $1, "password" = $2, "roleid" = $3 where "id" = $4`
	_, err = d.database.Exec(
		updateQuery,
		user.FullName,
		services.PasswordHash(user.Password),
		user.RoleID,
		user.ID)
	if err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (d *UserDB) UserDeleteOne(ID int32) error {
	u, err := d.UserFindOneId(ID)
	if u.ID == 0 {
		return errors.New("user not found")
	}
	deleteQuery := `delete from "Users" where "id" = $1`
	_, err = d.database.Exec(deleteQuery, ID)
	if err != sql.ErrNoRows {
		return err
	}
	return nil
}
