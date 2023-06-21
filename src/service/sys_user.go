package service

import (
	"crow-qin/src/config"
	"crow-qin/src/models"
	"github.com/toolkits/pkg/logger"
)

func UpdateUserInfo(updateUserInfoRequire UpdateUserInfoRequest) error {

	var sysUsers models.SysUser
	sysUsers.ID = updateUserInfoRequire.UserId
	sysUsers.Email = updateUserInfoRequire.Email
	sysUsers.Phone = updateUserInfoRequire.Phone
	logger.Info(sysUsers)
	err := models.UpdateUserInfoByID(sysUsers, config.MyTx)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func UpdateUserPassword(updateUserPasswordRequire UpdateUserPasswordRequest) (err error) {
	var sysUsers models.SysUser
	sysUsers.UserName = updateUserPasswordRequire.UserName
	sysUsers.Password = EncryptWithMd5(updateUserPasswordRequire.NewPassword)
	logger.Info(sysUsers)
	err = models.UpdateUserInfoByUserName(sysUsers, config.MyTx)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func GetUserInfo(userName string) (userInfo models.SysUser, err error) {
	userInfo, err = models.GetUserByUserName(userName, config.MyTx)
	if err != nil {
		return userInfo, err
	}
	logger.Info(userInfo)
	return userInfo, nil
}
