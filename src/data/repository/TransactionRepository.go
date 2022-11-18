package repository

import (
	"avito/src/data/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	transactionsTable = "transactions"
)

type TransactionRepository struct {
	DB *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (t *TransactionRepository) GetById(id int) (*entity.Transaction, error) {

	transaction := new(entity.Transaction)
	getTransactionQuery :=
		fmt.Sprintf("select * from %s as trans where trans.id = $1", transactionsTable)

	err := t.DB.Get(transaction, getTransactionQuery, id)
	return transaction, err
}

func (t *TransactionRepository) GetAll() ([]*entity.Transaction, error) {

	var transactions []*entity.Transaction
	getAllUsersQuery :=
		fmt.Sprintf("select * from %s", transactionsTable)

	if err := t.DB.Select(&transactions, getAllUsersQuery); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *TransactionRepository) Create(transaction *entity.Transaction) (*entity.Transaction, error) {

	tx, err := t.DB.Begin()
	if err != nil {
		return nil, err
	}

	var id int
	createTransactionQuery :=
		fmt.Sprintf("insert into %s (id, service_id, amount, from_id, to_id, status) values ($1, $2, $3, $4, $5, $6) returning id",
			transactionsTable)

	row := tx.QueryRow(createTransactionQuery, transaction.Id, transaction.ServiceId,
		transaction.Amount, transaction.FromId, transaction.ToId, transaction.TransactionStatus)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return nil, err
	}

	return transaction, tx.Commit()
}

func (t *TransactionRepository) Update(id int, transaction *entity.Transaction) (*entity.Transaction, error) {

	tx, err := t.DB.Begin()
	if err != nil {
		return nil, err
	}

	var foundId int
	getTransactionQuery :=
		fmt.Sprintf("select id from %s where id = $1", transactionsTable)

	row := tx.QueryRow(getTransactionQuery, id)
	if err := row.Scan(&foundId); err != nil {
		tx.Rollback()
		return t.Create(transaction)
	}

	updateTransactionQuery :=
		fmt.Sprintf("update %s set id = $1, amount = $2, from_id = $3, to_id = $4, status = $5 where id = $1", transactionsTable)

	if _, err := tx.Exec(updateTransactionQuery, transaction.Id, transaction.Amount,
		transaction.FromId, transaction.ToId, transaction.TransactionStatus); err != nil {
		tx.Rollback()
		return nil, err
	}

	return transaction, tx.Commit()
}

func (t *TransactionRepository) Delete(id int) (bool, error) {

	tx, err := t.DB.Begin()
	if err != nil {
		return false, err
	}

	deleteTransactionQuery :=
		fmt.Sprintf("delete from %s where id = $1",
			transactionsTable)

	if _, err := tx.Exec(deleteTransactionQuery, id); err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit()
}
