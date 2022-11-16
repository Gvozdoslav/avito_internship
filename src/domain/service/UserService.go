package service

import (
	"avito/src/data/entity"
	"avito/src/data/repository"
)

type UserService struct {
	users                 []entity.User
	userRepository        repository.UserRepository
	accountRepository     repository.AccountRepository
	transactionRepository repository.TransactionRepository
}
