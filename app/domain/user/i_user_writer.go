package user

import "ma-inject/app/domain/dtos/inputs"

type IUserWriter interface {
	IUserReader
	AddUser(user *inputs.UserInput) error
	UpdateUserNickName(id int64, nickName string) error
}
