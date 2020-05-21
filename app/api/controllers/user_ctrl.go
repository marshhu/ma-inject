package controllers

import (
	"github.com/gin-gonic/gin"
	"ma-inject/app/domain/dtos/inputs"
	"ma-inject/app/domain/user"
	"strconv"
)

type UserController struct {
	UserService user.IUserService `inject:"UserService"`
}

func (ctrl *UserController) GetUsers(ctx *gin.Context) {
	users := ctrl.UserService.GetUsers()
	Ok(Response{Code: Success, Msg: "获取用户成功！", Data: users}, ctx)
}

func (ctrl *UserController) GetUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		BadRequestError("id参数格式错误", ctx)
		return
	}
	users := ctrl.UserService.GetUser(id)
	Ok(Response{Code: Success, Msg: "获取用户成功！", Data: users}, ctx)
}

func (ctrl *UserController) AddUser(ctx *gin.Context) {
	input := inputs.UserInput{}
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		BadRequestError("参数错误", ctx)
		return
	}
	err = ctrl.UserService.AddUser(&input)
	if err != nil {
		Ok(Response{Code: Failed, Msg: err.Error()}, ctx)
		return
	}
	Ok(Response{Code: Success, Msg: "添加用户成功！"}, ctx)
}
