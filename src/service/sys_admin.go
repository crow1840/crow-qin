package service

import (
	"crow-qin/src/config"
	"crow-qin/src/models"
	"errors"
	"github.com/toolkits/pkg/logger"
	"regexp"
)

func CheckPassword(password string) (err error) {
	//校验密码长度
	if len(password) > 18 || len(password) < 6 {
		err = errors.New("密码必须是8~16位")
		return err
	} else {
		logger.Info(len(password), "长度检查通过")
	}

	//校验纯英文
	b, err := regexp.MatchString("^([A-z]+)$", password)
	if err != nil {
		logger.Error(err)
	}
	if b == true {
		err = errors.New("密码不能是纯英文")
		return err
	} else {
		logger.Info("纯英文检查通过")
	}

	//校验纯数字
	b, err = regexp.MatchString("^([0-9]+)$", password)
	if err != nil {
		logger.Error(err)
		return err
	}
	if b == true {
		err = errors.New("密码不能是纯数字")
		return err
	} else {
		logger.Info("纯数字检查通过")
	}
	return nil
}

func CreateUser(createUserRequire CreateUserRequest) (userId int64, err error) {

	//查重
	checkResult, err := models.CheckUserExists(createUserRequire.UserName, config.MyTx)
	if checkResult == true {
		err = errors.New("用户名已经存在")
		logger.Error(err)
		return 0, err
	} else {
		logger.Info("用户名查重通过")
	}

	//创建用户
	var sysUser models.SysUser
	sysUser.UserName = createUserRequire.UserName
	sysUser.Role = createUserRequire.Role
	sysUser.Phone = createUserRequire.Phone
	sysUser.Email = createUserRequire.Email
	sysUser.Password = EncryptWithMd5(createUserRequire.Password)

	userId, err = models.CreateUser(sysUser, config.MyTx)
	if err != nil {
		logger.Error(err)
		return userId, nil
	}
	return userId, nil

}

func ResetUserPassword(resetUserPasswordRequire ResetUserPasswordRequest) error {

	//检查用户是否存在
	updateUser, err := models.GetUserById(resetUserPasswordRequire.UserId, config.MyTx)
	if err != nil {
		logger.Error(err)
		return err
	}
	if len(updateUser.UserName) == 0 {
		err = errors.New("用户不存在")
		logger.Error(err)
		return err
	}

	//修改密码
	var sysUsers models.SysUser
	sysUsers.ID = resetUserPasswordRequire.UserId
	sysUsers.Password = EncryptWithMd5(resetUserPasswordRequire.NewPassword)
	logger.Info(sysUsers)
	err = models.UpdateUserInfoByID(sysUsers, config.MyTx)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil

}

func GetUsers(getUsersRequest GetUsersRequest) (getUsersResponse GetUsersResponse, err error) {
	//查询
	var sysUser models.SysUser
	sysUser.UserName = getUsersRequest.UserName
	sysUser.Role = getUsersRequest.Role
	sysUsers, count, err := models.GetUsersPage(getUsersRequest.PageSize, getUsersRequest.PageNum, sysUser, config.MyTx)
	if err != nil {
		logger.Error(err)
		return getUsersResponse, err
	}
	//组装返回结果
	getUsersResponse.SysUsers = sysUsers
	getUsersResponse.PageSize = getUsersRequest.PageSize
	getUsersResponse.PageNum = getUsersRequest.PageNum
	getUsersResponse.Count = count

	return getUsersResponse, nil
}

func GetUser(userId int64) (sysUser models.SysUser, err error) {
	sysUser, err = models.GetUserById(userId, config.MyTx)
	if err != nil {
		logger.Error(err)
		return sysUser, err
	}
	return sysUser, nil
}

func DeleteUser(userId int64) error {
	//检查用户是否存在
	updateUser, err := models.GetUserById(userId, config.MyTx)
	if err != nil {
		logger.Error(err)
		return err
	}
	if len(updateUser.UserName) == 0 {
		err = errors.New("用户不存在")
		logger.Error(err)
		return err
	}

	//删除用户
	err = models.DeleteUserById(userId, config.MyTx)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func UpdateUsersInfo(updateUserInfoRequire UpdateUsersInfoRequest) (err error) {
	var sysUsers models.SysUser
	sysUsers.ID = updateUserInfoRequire.UserId
	sysUsers.Email = updateUserInfoRequire.Email
	sysUsers.Phone = updateUserInfoRequire.Phone
	logger.Info(sysUsers)
	err = models.UpdateUserInfoByID(sysUsers, config.MyTx)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
