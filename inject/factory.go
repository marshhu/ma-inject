package inject

import (
	"ma-inject/app/api/controllers"
)

type CtrlFactory struct {
	UserCtrl *controllers.UserController `inject:"UserController"`
}
