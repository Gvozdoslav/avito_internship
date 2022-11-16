package service

import (
	"avito/src/domain/dto"
)

type UserServiceInterface interface {
	getUser(id int) dto.UserDto
	getAllUsers() []dto.UserDto
	createUser() dto.UserDto
	updateUser(id int) dto.UserDto
	deleteUser(id int) bool
}
