package service

import (
	"avito/src/data/entity"
	"avito/src/data/repository"
	"avito/src/domain/dto"
)

type UserService struct {
	userRepository        repository.Repository[entity.User]
	accountRepository     repository.Repository[entity.Account]
	transactionRepository repository.Repository[entity.Transaction]
}

func NewUserService(
	userRepository repository.Repository[entity.User],
	accountRepository repository.Repository[entity.Account],
	transactionRepository repository.Repository[entity.Transaction]) *UserService {

	return &UserService{
		userRepository:        userRepository,
		accountRepository:     accountRepository,
		transactionRepository: transactionRepository,
	}
}

func (u UserService) getUser(id int) dto.UserDto {
	//TODO implement me
	panic("implement me")
}

func (u UserService) getAllUsers() []dto.UserDto {
	//TODO implement me
	panic("implement me")
}

func (u UserService) createUser() dto.UserDto {
	//TODO implement me
	panic("implement me")
}

func (u UserService) updateUser(id int) dto.UserDto {
	//TODO implement me
	panic("implement me")
}

func (u UserService) deleteUser(id int) bool {
	//TODO implement me
	panic("implement me")
}
