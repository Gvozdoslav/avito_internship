package service

import (
	"avito/src/data/dto"
)

type UserServiceInterface interface {
	GetUser(id int) (*dto.UserDto, error)
	GetAllUsers() ([]*dto.UserDto, error)
	CreateUser(userDto *dto.UserDto) (*dto.UserDto, error)
	UpdateUser(id int) (*dto.UserDto, error)
	DeleteUser(id int) (bool, error)
}
