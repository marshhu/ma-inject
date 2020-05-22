package user

import (
	"ma-inject/app/domain/dtos"
	"ma-inject/app/domain/dtos/inputs"
)

type UserService struct {
	UserRepository IUserRepository `inject:"UserRepository"`
}

func (s *UserService) AddUser(user *inputs.UserInput) error {
	return s.UserRepository.AddUser(user)
}

func (s *UserService) GetUsers() []dtos.UserDto {
	return s.UserRepository.GetUsers()
}

func (s *UserService) GetUser(id int64) *dtos.UserDto {
	return s.UserRepository.GetUser(id)
}
