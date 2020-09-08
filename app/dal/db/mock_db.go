package db

import "github.com/marshhu/ma-inject/app/dal/entities"

//准备用户数据，实际开发一般从数据库读取
var users []entities.UserEntity

func init() {
	users = append(users, entities.UserEntity{ID: 1, Name: "小明", NickName: "无敌", Gender: 1, Age: 13, Tel: "18886588086", Address: "中国，广东，深圳"})
	users = append(users, entities.UserEntity{ID: 2, Name: "小红", NickName: "傻妞", Gender: 0, Age: 13, Tel: "1888658809", Address: "中国，广东，广州"})
}

type MockDB struct {
	Host  string
	User  string
	Pwd   string
	Alias string
}

func (db *MockDB) Connect() bool {
	return true
}

func (db *MockDB) Users() []entities.UserEntity {
	return users
}

func (db *MockDB) Close() {

}
