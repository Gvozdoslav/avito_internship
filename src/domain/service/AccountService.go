package service

import (
	"avito/src/data/dto"
	"avito/src/data/repository"
)

type AccountService struct {
	accountRepository *repository.AccountRepository
}

func NewAccountService(accountRepository *repository.AccountRepository) *AccountService {
	return &AccountService{
		accountRepository: accountRepository,
	}
}

func (a *AccountService) GetAccount(id int) (*dto.AccountDto, error) {

	account, err := a.accountRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	accountResponseDto := dto.ToAccountDto(account)
	return accountResponseDto, nil
}

func (a *AccountService) GetAllAccounts() ([]*dto.AccountDto, error) {

	accounts, err := a.accountRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var accountsResponseDto []*dto.AccountDto

	for _, elem := range accounts {
		accountsResponseDto = append(accountsResponseDto, dto.ToAccountDto(elem))
	}

	return accountsResponseDto, nil
}

func (a *AccountService) CreateAccount(accountDto *dto.AccountDto) (*dto.AccountDto, error) {

	account, err := a.accountRepository.Create(accountDto.ToAccountEntity())
	if err != nil {
		return nil, err
	}

	accountResponseDto := dto.ToAccountDto(account)
	return accountResponseDto, nil
}

func (a *AccountService) UpdateAccount(id int, accountDto *dto.AccountDto) (*dto.AccountDto, error) {

	accountEntity := accountDto.ToAccountEntity()

	account, err := a.accountRepository.Update(id, accountEntity)
	if err != nil {
		return nil, err
	}

	accountResponseDto := dto.ToAccountDto(account)
	return accountResponseDto, nil
}

func (a *AccountService) Delete(id int) (bool, error) {

	isOk, err := a.accountRepository.Delete(id)
	if err != nil {
		return isOk, err
	}

	return true, nil
}
