package user

import (
	"ma-inject/app/domain/dtos"
	"ma-inject/app/domain/dtos/inputs"
)

type IUserService interface {
	GetUsers() []dtos.UserDto
	GetUser(id int64) *dtos.UserDto
	AddUser(user *inputs.UserInput) error
}
