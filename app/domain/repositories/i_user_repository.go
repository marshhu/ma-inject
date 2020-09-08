package repositories

import (
	"github.com/marshhu/ma-inject/app/domain/dtos/inputs"
)

type IUserRepository interface {
	IUserReader
	AddUser(user *inputs.UserInput) error
	UpdateUserNickName(id int64, nickName string) error
}
