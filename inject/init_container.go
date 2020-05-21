package inject

import (
	"fmt"
	"ma-inject/app/api/controllers"
	"ma-inject/app/dal/db"
	"ma-inject/app/dal/user"
	userDomain "ma-inject/app/domain/user"
)

var GContainer = &Container{
	singletons: make(map[string]interface{}),
	factories:  make(map[string]factory),
}

func Init() {
	//db
	GContainer.SetSingleton("MockDBRead", &db.MockDB{Host: "192.168.1.12:3036", User: "root", Pwd: "123456", Alias: "Read"})
	GContainer.SetSingleton("MockDBWrite", &db.MockDB{Host: "192.168.1.25:3036", User: "root", Pwd: "123456", Alias: "Write"})

	//仓储
	GContainer.SetSingleton("UserRead", &user.UserRead{})
	GContainer.SetSingleton("UserWrite", &user.UserWrite{})

	//服务
	GContainer.SetSingleton("UserService", &userDomain.UserService{})

	//控制器
	GContainer.SetSingleton("UserController", &controllers.UserController{})

	//控制器工厂
	ctlFactory := &CtrlFactory{}
	GContainer.SetSingleton("CtrlFactory", ctlFactory)

	GContainer.Entry(ctlFactory) //注入

	fmt.Println(GContainer.String())
}
