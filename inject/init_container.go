package inject

import (
	"fmt"
	"github.com/marshhu/ma-inject/app/api/controllers"
	"github.com/marshhu/ma-inject/app/dal/db"
	"github.com/marshhu/ma-inject/app/dal/user"
	userDomain "github.com/marshhu/ma-inject/app/domain/user"
)

var GContainer = &Container{
	singletons: make(map[string]interface{}),
	transients: make(map[string]factory),
}

func Init() {
	//db
	GContainer.SetSingleton("MockDBRead", &db.MockDB{Host: "192.168.1.12:3036", User: "root", Pwd: "123456", Alias: "Read"})
	GContainer.SetSingleton("MockDBWrite", &db.MockDB{Host: "192.168.1.25:3036", User: "root", Pwd: "123456", Alias: "Write"})

	//仓储
	GContainer.SetSingleton("UserRepository", &user.UserRepository{})

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
