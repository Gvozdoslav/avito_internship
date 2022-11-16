package repository

import (
	"avito/src/data/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u UserRepository) getById(id int) entity.User {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) getAll() entity.User {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) create() entity.User {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) update(id int) entity.User {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) delete(id int) entity.User {
	//TODO implement me
	panic("implement me")
}
