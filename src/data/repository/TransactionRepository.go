package repository

import (
	"avito/src/data/entity"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	DB *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (t TransactionRepository) getById(id int) entity.Transaction {
	//TODO implement me
	panic("implement me")
}

func (t TransactionRepository) getAll() entity.Transaction {
	//TODO implement me
	panic("implement me")
}

func (t TransactionRepository) create() entity.Transaction {
	//TODO implement me
	panic("implement me")
}

func (t TransactionRepository) update(id int) entity.Transaction {
	//TODO implement me
	panic("implement me")
}

func (t TransactionRepository) delete(id int) entity.Transaction {
	//TODO implement me
	panic("implement me")
}
