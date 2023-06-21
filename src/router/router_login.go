package router

import (
	"crow-qin/src/service"
	"github.com/gin-gonic/gin"
)

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录，并获取token
// @Tags system
// @Accept  application/json
// @Product application/json
// @Param data body service.LoginRequest  true "用户名，用户密码"
// @Success 200 {object} response.Response{data=string} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/login [post]
// @Security Bearer
func Login(c *gin.Context) {

	var loginRequest service.LoginRequest
	err := c.Bind(&loginRequest)
	if err != nil {
		SetErrDefault(c, err)
		return
	}
	signedToken, err := service.Login(loginRequest)
	if err != nil {
		SetErrDefault(c, err)
		return
	}
	SetOK(c, signedToken)
}

// Refresh 刷新token
// @Summary 刷新token
// @Description 刷新用户token
// @Tags system
// @Accept  application/json
// @Product application/json
// @Success 200 {object} response.Response{data=string} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/login/refresh [post]
// @Security Bearer
func Refresh(c *gin.Context) {
	strToken, err := service.GetTokenInHeader(c)
	if err != nil {
		SetErrDefault(c, err)
		return
	}
	newToken, err := service.RefreshToken(strToken)
	if err != nil {
		SetErrDefault(c, err)
		return
	}
	SetOK(c, newToken)
}

// Logout 用户登出
// @Summary 用户登出
// @Description 用户登出
// @Tags system
// @Accept  application/json
// @Product application/json
// @Success 200 {object} response.Response{data=string} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/logout [post]
// @Security Bearer
func Logout(c *gin.Context) {

	//从header中获取token
	strToken, err := service.GetTokenInHeader(c)
	if err != nil {
		SetErrDefault(c, err)
		return
	}
	//删除redis中的token
	err = service.Logout(strToken)
	if err != nil {
		SetErrDefault(c, err)
		return
	}
	SetOK(c, "登出成功")
}
