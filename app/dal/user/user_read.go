package user

import (
	"github.com/marshhu/ma-inject/app/dal/db"
	"github.com/marshhu/ma-inject/app/domain/dtos"
)

type UserRead struct {
	ReadDb *db.MockDB `inject:"MockDBRead"`
}

func (r *UserRead) GetUsers() []dtos.UserDto {
	if r.ReadDb.Connect() {
		users := r.ReadDb.Users()
		var list []dtos.UserDto
		for _, user := range users {
			list = append(list, dtos.UserDto{ID: user.ID, Name: user.Name, NickName: user.NickName, Gender: user.Gender, Age: user.Age, Tel: user.Tel, Address: user.Address})
		}
		return list
	}
	return nil
}

func (r *UserRead) GetUser(id int64) *dtos.UserDto {
	if r.ReadDb.Connect() {
		users := r.ReadDb.Users()
		for _, user := range users {
			if user.ID == id {
				return &dtos.UserDto{ID: user.ID, Name: user.Name, NickName: user.NickName, Gender: user.Gender, Age: user.Age, Tel: user.Tel, Address: user.Address}
			}
		}
		return &dtos.UserDto{}
	}
	return nil
}

func (r *UserRead) GetMaxUserId() int64 {
	var maxId int64
	if r.ReadDb.Connect() {
		users := r.ReadDb.Users()
		for _, user := range users {
			if user.ID > maxId {
				maxId = user.ID
			}
		}
	}
	return maxId
}
