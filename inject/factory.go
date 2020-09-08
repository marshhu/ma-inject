package inject

import (
	"github.com/marshhu/ma-inject/app/api/controllers"
)

type CtrlFactory struct {
	UserCtrl *controllers.UserController `inject:"UserController"`
}
