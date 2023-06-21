package router

import (
	"crow-qin/src/config"
	"crow-qin/src/models"
	"crow-qin/src/service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/toolkits/pkg/logger"
)

// UpdateUserInfo 用户修改自己信息
// @Summary 用户修改自己信息
// @Description 用户修改自己信息
// @Tags user
// @Accept  application/json
// @Product application/json
// @Param data body UpdateUserInfoRequest  true "选项可选"
// @Success 200 {object} response.Response{data=string} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/user/info [put]
// @Security Bearer
func UpdateUserInfo(c *gin.Context) {
	var updateUserInfoRequire UpdateUserInfoRequest
	err := c.Bind(&updateUserInfoRequire)
	if err != nil {
		SetErrDefault(c, err)
		return
	}

	userIdAny, _ := c.Get("user_id")
	//更新数据库
	var updateUserInfoRequireS service.UpdateUserInfoRequest
	updateUserInfoRequireS.UserId = userIdAny.(int64)
	updateUserInfoRequireS.Email = updateUserInfoRequire.Email
	updateUserInfoRequireS.Phone = updateUserInfoRequire.Phone
	err = service.UpdateUserInfo(updateUserInfoRequireS)
	SetOK(c, "用户信息修改成功")
}

func user() gin.HandlerFunc {

	return func(c *gin.Context) {
		idAuth, myName, err := service.Verify(c)
		if err != nil {
			SetErrDefault(c, err)
			//c.Abort()
			return
		}
		if idAuth == false {
			err = errors.New("身份认证失败")
			SetErrDefault(c, err)
			//c.Abort()
			return
		}
		userInfo, err := models.GetUserByUserName(myName, config.MyTx)
		if err != nil {
			SetErrDefault(c, err)
			return
		}
		logger.Info(userInfo)

		c.Set("user_role", userInfo.Role)
		c.Set("user_name", userInfo.UserName)
		c.Set("user_id", userInfo.ID)
		c.Next()
	}

}

// UpdateUserPassword 用户修改密码
// @Summary 用户修改密码
// @Description 用户修改自己的密码
// @Tags user
// @Accept  application/json
// @Product application/json
// @Param data body UpdateUserPasswordRequest  true "用户的新密码和旧密码"
// @Success 200 {object} response.Response{data=string} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/user/password [put]
// @Security Bearer
func UpdateUserPassword(c *gin.Context) {
	var updateUserPasswordRequire UpdateUserPasswordRequest

	err := c.Bind(&updateUserPasswordRequire)
	if err != nil {
		SetErrDefault(c, err)
		return
	}

	userIdAny, _ := c.Get("user_name")
	userName := userIdAny.(string)

	//旧密码认证
	var loginRequest service.LoginRequest
	loginRequest.UserName = userName
	loginRequest.Password = updateUserPasswordRequire.Password
	userId, err := service.CheckUserPassword(loginRequest)
	if err != nil {
		SetErrDefault(c, err)
		return
	}
	logger.Info(userId)

	var updateUserPasswordRequireS service.UpdateUserPasswordRequest
	err = service.UpdateUserPassword(updateUserPasswordRequireS)

	SetOK(c, "密码修改成功")
}

// GetUserInfo 用户获取本人信息
// @Summary 用户获取本人信息
// @Description 用户获取本人信息
// @Tags user
// @Success 200 {object} response.Response{data=models.SysUser} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/user/info [get]
// @Security Bearer
func GetUserInfo(c *gin.Context) {

	userNameAny, b := c.Get("user_name")
	if b == false {
		err := errors.New("用户认证失败")
		SetErrDefault(c, err)
		return
	}
	userName := userNameAny.(string)
	//查询信息
	userInfo, err := service.GetUserInfo(userName)
	if err != nil {
		SetErrDefault(c, err)
		return
	}
	SetOK(c, userInfo)

}
