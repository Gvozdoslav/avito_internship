package repository

import (
	"avito/src/data/entity"
	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	DB *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

func (a AccountRepository) getById(id int) entity.Account {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepository) getAll() entity.Account {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepository) create() entity.Account {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepository) update(id int) entity.Account {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepository) delete(id int) entity.Account {
	//TODO implement me
	panic("implement me")
}
