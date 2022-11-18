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

func (a *AccountService) BookForService(sendId int, recId int, amount float64) error {

	sender, err := a.GetAccount(sendId)
	if err != nil || sender == nil {
		return err
	}

	receiver, err := a.GetAccount(recId)
	if err != nil || receiver == nil {
		return err
	}

	if sender.Balance < amount {
		return err
	}

	return nil
}

func (a *AccountService) PayForService(sendId int, recId int, amount float64) error {

	sender, err := a.GetAccount(sendId)
	if err != nil || sender == nil {
		return err
	}

	receiver, err := a.GetAccount(recId)
	if err != nil || receiver == nil {
		return err
	}

	if sender.Balance < amount {
		return err
	}

	sender.Balance -= amount
	receiver.Balance += amount

	_, err = a.UpdateAccount(sendId, sender)
	if err != nil {
		return err
	}

	_, err = a.UpdateAccount(recId, receiver)
	if err != nil {
		return err
	}

	return nil
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
