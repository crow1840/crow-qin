package router

import (
	"crow-qin/src/config"
	"crow-qin/src/models"
	"crow-qin/src/service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/toolkits/pkg/logger"
)

// CreateUser 创建用户
// @Summary 创建用户
// @Description 创建用户
// @Tags admin
// @Accept  application/json
// @Product application/json
// @Param data body service.CreateUserRequest  true "填写用户信息"
// @Success 200 {object} response.Response{data=int64} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/admin/user [post]
// @Security Bearer
func CreateUser(c *gin.Context) {

	var createUserRequire service.CreateUserRequest

	err := c.Bind(&createUserRequire)
	if err != nil {
		SetErrDefault(c, err)
		return
	}

	//输入检查
	if createUserRequire.UserName == "" {
		err = errors.New("用户名不能为空")
		SetErrDefault(c, err)
		return
	}
	if createUserRequire.Role == "" {
		err = errors.New("用户角色不能为空")
		SetErrDefault(c, err)
		return
	}
	err = service.CheckPassword(createUserRequire.Password)
	if err != nil {
		SetErrDefault(c, err)
		return
	}

	//查重
	userId, err := service.CreateUser(createUserRequire)
	if err != nil {
		SetErrDefault(c, err)
		return
	}
	SetOK(c, userId)

}

func admin() gin.HandlerFunc {
	return func(c *gin.Context) {

		//从token判断
		idAuth, myName, err := service.Verify(c) //第二个值是用户名，这里没有使用
		if err != nil {
			SetErrDefault(c, err)
			return
		}
		if idAuth == false {
			err = errors.New("身份认证失败")
			SetErrDefault(c, err)
			return
		}
		myRole, err := models.GetRoleByUserName(myName, config.MyTx)
		logger.Info(myRole)
		if myRole != "admin" {
			err = errors.New("没有管理员权限")
			SetErrDefault(c, err)
			return
		}
		c.Next()
	}
}

// ResetUserPassword admin重置用户密码
// @Summary admin重置用户密码
// @Description admin重置用户密码
// @Tags admin
// @Accept  application/json
// @Product application/json
// @Param data body service.ResetUserPasswordRequest  true "填写用户Id和xin密码"
// @Success 200 {object} response.Response{data=string} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/admin/user-password [put]
// @Security Bearer
func ResetUserPassword(c *gin.Context) {
	var resetUserPasswordRequire service.ResetUserPasswordRequest
	err := c.Bind(&resetUserPasswordRequire)
	if err != nil {
		SetErrDefault(c, err)
		return
	}

	//输入检查
	if &resetUserPasswordRequire.UserId == nil {
		err = errors.New("用户id不能为空")
		SetErrDefault(c, err)
		return
	}
	err = service.CheckPassword(resetUserPasswordRequire.NewPassword)
	if err != nil {
		SetErrDefault(c, err)
		return
	}

	//修改密码
	err = service.ResetUserPassword(resetUserPasswordRequire)
	if err != nil {
		SetErrDefault(c, err)
		return
	}
	SetOK(c, "密码重置成功")
}

// GetUsers 获取用户信息
// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags admin
// @Param role query string false "用户角色"
// @Param user_name query string false "用户名（可模糊查询）"
// @Param page_num query string false "页数"
// @Param page_size query string false "每页行数"
// @Success 200 {object} response.Response{data=service.GetUsersResponse} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/admin/users [get]
// @Security Bearer
func GetUsers(c *gin.Context) {

	var getUsersRequest service.GetUsersRequest
	pageNumString := c.Query("page_num")
	getUsersRequest.PageNum = service.StringToInt64(pageNumString)
	pageSizeString := c.Query("page_size")
	getUsersRequest.PageSize = service.StringToInt64(pageSizeString)
	getUsersRequest.Role = c.Query("role")
	getUsersRequest.UserName = c.Query("user_name")

	//查询
	getUsersResponse, err := service.GetUsers(getUsersRequest)
	if err != nil {
		SetErrDefault(c, err)
		return
	}
	SetOK(c, getUsersResponse)

}

// GetUser 获取指定用户信息
// @Summary 获取指定用户信息
// @Description 获取指定用户信息
// @Tags admin
// @Param uuid path string true "用户id"
// @Success 200 {object} response.Response{data=models.SysUser} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/admin/users/{uuid} [get]
// @Security Bearer
func GetUser(c *gin.Context) {
	uuid := c.Param("uuid")
	userId := service.StringToInt64(uuid)
	//查询
	var sysUser models.SysUser

	sysUser, err := service.GetUser(userId)
	if err != nil {
		SetErrDefault(c, err)
		return
	}
	SetOK(c, sysUser)

}

// DeleteUser 删除指定用户
// @Summary 删除指定用户
// @Description 删除指定用户
// @Tags admin
// @Param uuid path string true "用户id"
// @Success 200 {object} response.Response{data=string} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/admin/users/{uuid} [delete]
// @Security Bearer
func DeleteUser(c *gin.Context) {
	uuid := c.Param("uuid")
	userId := service.StringToInt64(uuid)
	//检查用户是否存在

	err := service.DeleteUser(userId)
	if err != nil {
		SetErrDefault(c, err)
		return
	}
	SetOK(c, "删除成功")
}

// UpdateUsersInfo 修改用户信息
// @Summary 修改用户信息
// @Description 修改用户信息
// @Tags admin
// @Accept  application/json
// @Product application/json
// @Param data body service.UpdateUsersInfoRequest  true "user_id必须，其他可选"
// @Success 200 {object} response.Response{data=string} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/admin/users [put]
// @Security Bearer
func UpdateUsersInfo(c *gin.Context) {
	var updateUserInfoRequire service.UpdateUsersInfoRequest
	err := c.Bind(&updateUserInfoRequire)
	if err != nil {
		SetErrDefault(c, err)
		return
	}

	//更新数据库
	err = service.UpdateUsersInfo(updateUserInfoRequire)
	if err != nil {
		SetErrDefault(c, err)
		return
	}
	SetOK(c, "用户信息修改成功")
}
