package repository

import (
	"avito/src/data/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) GetById(id int) (*entity.User, error) {

	user := new(entity.User)
	getUserQuery :=
		fmt.Sprintf("select * from %s as usrs where usrs.id = $1", usersTable)

	err := u.DB.Get(user, getUserQuery, id)
	return user, err
}

func (u UserRepository) GetAll() ([]*entity.User, error) {

	var users []*entity.User
	getAllUsersQuery :=
		fmt.Sprintf("select * from %s", usersTable)

	if err := u.DB.Select(&users, getAllUsersQuery); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) Create(user *entity.User) (*entity.User, error) {

	tx, err := u.DB.Begin()
	if err != nil {
		return nil, err
	}

	var id int
	createUserQuery :=
		fmt.Sprintf("insert into %s (id, name, username, password, account_id) values ($1, $2, $3, $4, $5) returning id",
			usersTable)

	row := tx.QueryRow(createUserQuery, user.Id, user.Name, user.Username, user.Password, user.Account)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return nil, err
	}

	return user, tx.Commit()
}

func (u *UserRepository) Update(id int, user *entity.User) (*entity.User, error) {

	tx, err := u.DB.Begin()
	if err != nil {
		return nil, err
	}

	var foundId int
	getUserQuery :=
		fmt.Sprintf("select id from %s where id = $1", usersTable)

	row := tx.QueryRow(getUserQuery, id)
	if err := row.Scan(&foundId); err != nil {
		tx.Rollback()
		return u.Create(user)
	}

	updateUserQuery :=
		fmt.Sprintf("update %s set id = $1, name = $2, username = $3, password = $4 where id = $1", usersTable)

	if _, err := tx.Exec(updateUserQuery, user.Id, user.Name, user.Username, user.Password); err != nil {
		tx.Rollback()
		return nil, err
	}

	return user, tx.Commit()
}

func (u *UserRepository) Delete(id int) (bool, error) {

	tx, err := u.DB.Begin()
	if err != nil {
		return false, err
	}

	deleteQuery :=
		fmt.Sprintf("delete from %s where id = $1",
			usersTable)

	if _, err := tx.Exec(deleteQuery, id); err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit()
}
