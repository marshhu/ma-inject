package user

import (
	"ma-inject/app/domain/dtos"
	"ma-inject/app/domain/dtos/inputs"
)

type UserService struct {
	UserWriter IUserWriter `inject:"UserWrite"`
}

func (s *UserService) AddUser(user *inputs.UserInput) error {
	return s.UserWriter.AddUser(user)
}

func (s *UserService) GetUsers() []dtos.UserDto {
	return s.UserWriter.GetUsers()
}

func (s *UserService) GetUser(id int64) *dtos.UserDto {
	return s.UserWriter.GetUser(id)
}
