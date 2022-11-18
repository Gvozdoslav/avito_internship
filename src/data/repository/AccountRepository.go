package repository

import (
	"avito/src/data/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	accountsTable = "accounts"
)

type AccountRepository struct {
	DB *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

func (a *AccountRepository) GetById(id int) (*entity.Account, error) {

	account := new(entity.Account)
	getAccountQuery :=
		fmt.Sprintf("select * from %s as accs where accs.id = $1", accountsTable)

	err := a.DB.Get(account, getAccountQuery, id)
	return account, err
}

func (a *AccountRepository) GetAll() ([]*entity.Account, error) {

	var accounts []*entity.Account
	getAllUsersQuery :=
		fmt.Sprintf("select * from %s", accountsTable)

	if err := a.DB.Select(&accounts, getAllUsersQuery); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (a *AccountRepository) Create(account *entity.Account) (*entity.Account, error) {

	tx, err := a.DB.Begin()
	if err != nil {
		return nil, err
	}

	var id int
	createAccountQuery :=
		fmt.Sprintf("insert into %s (id, balance) values ($1, $2) returning id",
			accountsTable)

	row := tx.QueryRow(createAccountQuery, account.Id, account.Balance)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return nil, err
	}

	return account, tx.Commit()
}

func (a *AccountRepository) Update(id int, account *entity.Account) (*entity.Account, error) {

	tx, err := a.DB.Begin()
	if err != nil {
		return nil, err
	}

	var foundId int
	getAccountQuery :=
		fmt.Sprintf("select id from %s where id = $1", accountsTable)

	row := tx.QueryRow(getAccountQuery, id)
	if err := row.Scan(&foundId); err != nil {
		tx.Rollback()
		return a.Create(account)
	}

	updateAccountQuery :=
		fmt.Sprintf("update %s set id = $1, balance = $2 where id = $1", accountsTable)

	if _, err := tx.Exec(updateAccountQuery, account.Id, account.Balance); err != nil {
		tx.Rollback()
		return nil, err
	}

	return account, tx.Commit()
}

func (a *AccountRepository) Delete(id int) (bool, error) {

	tx, err := a.DB.Begin()
	if err != nil {
		return false, err
	}

	deleteAccountQuery :=
		fmt.Sprintf("delete from %s where id = $1",
			accountsTable)

	if _, err := tx.Exec(deleteAccountQuery, id); err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit()
}
