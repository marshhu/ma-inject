package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Success = 1  //成功状态码
	Failed  = -1 //失败状态码
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PageView struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

func Ok(r Response, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, r)
}

func BadRequestError(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, msg)
}

func InternalServerError(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, msg)
}

func Unauthorized(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, msg)
}

func Forbidden(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusForbidden, msg)
}

func NotFound(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, msg)
}

func NoContent(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, msg)
}
