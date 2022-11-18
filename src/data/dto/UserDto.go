package dto

import "avito/src/data/entity"

type UserDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Account  int    `json:"account"`
}

func (userdto *UserDto) ToUserEntity() *entity.User {

	return &entity.User{
		Id:       userdto.Id,
		Name:     userdto.Name,
		Username: userdto.Username,
		Password: userdto.Password,
		Account:  userdto.Account,
	}
}

func ToUserDto(u *entity.User) *UserDto {

	return &UserDto{
		Id:       u.Id,
		Name:     u.Name,
		Username: u.Username,
		Password: u.Password,
		Account:  u.Account,
	}
}
