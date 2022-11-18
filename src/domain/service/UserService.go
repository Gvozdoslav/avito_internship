package service

import (
	"avito/src/data/dto"
	"avito/src/data/repository"
)

type UserService struct {
	userRepository        *repository.UserRepository
	accountRepository     *repository.AccountRepository
	transactionRepository *repository.TransactionRepository
}

func NewUserService(
	userRepository *repository.UserRepository,
	accountRepository *repository.AccountRepository,
	transactionRepository *repository.TransactionRepository) *UserService {

	return &UserService{
		userRepository:        userRepository,
		accountRepository:     accountRepository,
		transactionRepository: transactionRepository,
	}
}

func (u *UserService) GetUser(id int) (*dto.UserDto, error) {

	user, err := u.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	userResponseDto := dto.ToUserDto(user)
	return userResponseDto, nil
}

func (u *UserService) GetAllUsers() ([]*dto.UserDto, error) {

	users, err := u.userRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var usersResponseDto []*dto.UserDto

	for _, elem := range users {
		usersResponseDto = append(usersResponseDto, dto.ToUserDto(elem))
	}

	return usersResponseDto, nil
}

func (u *UserService) CreateUser(userDto *dto.UserDto) (*dto.UserDto, error) {

	user, err := u.userRepository.Create(userDto.ToUserEntity())
	if err != nil {
		return nil, err
	}

	userResponseDto := dto.ToUserDto(user)
	return userResponseDto, nil
}

func (u *UserService) UpdateUser(id int, userDto *dto.UserDto) (*dto.UserDto, error) {

	userEntity := userDto.ToUserEntity()

	user, err := u.userRepository.Update(id, userEntity)
	if err != nil {
		return nil, err
	}

	userResponseDto := dto.ToUserDto(user)
	return userResponseDto, nil
}

func (u *UserService) DeleteUser(id int) (bool, error) {

	isOk, err := u.userRepository.Delete(id)
	if err != nil {
		return isOk, err
	}

	return true, nil
}
