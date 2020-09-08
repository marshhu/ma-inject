package user

import (
	"github.com/marshhu/ma-inject/app/domain/dtos"
	"github.com/marshhu/ma-inject/app/domain/dtos/inputs"
)

type IUserService interface {
	GetUsers() []dtos.UserDto
	GetUser(id int64) *dtos.UserDto
	AddUser(user *inputs.UserInput) error
}
