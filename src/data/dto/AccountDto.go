package dto

import "avito/src/data/entity"

type AccountDto struct {
	Id      int     `json:"id"`
	Balance float64 `json:"balance"`
}

func (accountDto *AccountDto) ToAccountEntity() *entity.Account {

	return &entity.Account{
		Id:      accountDto.Id,
		Balance: accountDto.Balance,
	}
}

func ToAccountDto(a *entity.Account) *AccountDto {

	return &AccountDto{
		Id:      a.Id,
		Balance: a.Balance,
	}
}
