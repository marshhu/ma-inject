package dtos

type UserDto struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nick_name"`
	Gender   int    `json:"gender"`
	Age      int    `json:"age"`
	Tel      string `json:"tel"`
	Address  string `json:"address"`
}
