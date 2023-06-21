package models

import (
	"errors"
	"github.com/toolkits/pkg/logger"
	"gorm.io/gorm"
)

type SysUser struct {
	*gorm.Model
	ID       int64  `json:"ID"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

func CheckUser(userName string, passWord string, tx *gorm.DB) (userId int64, err error) {
	var userInfo SysUser
	var count int64
	err = tx.Where("user_name = ? and password = ? ", userName, passWord).Find(&userInfo).Count(&count).Error
	if err != nil {
		return userId, err
	}
	if count < 0 {
		err = errors.New("未查询到用户")
		return userId, err
	}
	return userInfo.ID, nil
}

func CreateUser(sysUser SysUser, tx *gorm.DB) (userId int64, err error) {
	err = tx.Save(&sysUser).Error
	if err != nil {
		return userId, err
	}
	return sysUser.ID, nil
}

func CheckUserExists(userName string, tx *gorm.DB) (checkResult bool, err error) {
	var count int64
	err = tx.Model(&SysUser{}).Where("user_name = ?", userName).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		checkResult = true
	} else {
		checkResult = false
	}
	return checkResult, nil
}

func UpdateUserInfoByID(sysUsers SysUser, tx *gorm.DB) (err error) {
	err = tx.Where("id = ?", sysUsers.ID).Updates(&sysUsers).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateUserInfoByUserName(sysUsers SysUser, tx *gorm.DB) (err error) {
	err = tx.Where("user_name = ?", sysUsers.UserName).Updates(&sysUsers).Error
	if err != nil {
		return err
	}
	return nil
}

func GetRoleByUserName(userName string, tx *gorm.DB) (role string, err error) {
	err = tx.Model(&SysUser{}).Where("user_name = ?", userName).Pluck("role", &role).Error
	if err != nil {
		return role, err
	}
	return role, nil
}

func GetUserById(id int64, tx *gorm.DB) (sysUser SysUser, err error) {

	err = tx.Model(&SysUser{}).Where("id = ?", id).Find(&sysUser).Error
	if err != nil {
		return sysUser, err
	}
	return sysUser, nil
}

func GetUserByUserName(userName string, tx *gorm.DB) (sysUser SysUser, err error) {

	err = tx.Model(&SysUser{}).Where("user_name = ?", userName).Find(&sysUser).Error
	if err != nil {
		return sysUser, err
	}
	return sysUser, nil
}

func GetUsers(tx *gorm.DB) (sysUsers []SysUser, err error) {

	err = tx.Model(&SysUser{}).Find(&sysUsers).Error
	if err != nil {
		return sysUsers, err
	}
	return sysUsers, nil
}

func GetUsersPageB(pageSize int64, pageNum int64, sysUser SysUser, tx *gorm.DB) (sysUsers []SysUser, count int64, err error) {
	logger.Info(sysUser)
	err = tx.Model(sysUser).Where(sysUser).Count(&count).Error
	if err != nil {
		logger.Error(err)
	}
	if sysUser.UserName != "" {
		sysUser.UserName = "\\%" + sysUser.UserName + "\\%"
	}
	err = tx.Where(sysUser).Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&sysUsers).Error
	if err != nil {
		logger.Error(err)
		return sysUsers, count, err
	}
	return sysUsers, count, err
}

func GetUsersPage(pageSize int64, pageNum int64, sysUser SysUser, tx *gorm.DB) (sysUsers []SysUser, count int64, err error) {
	logger.Info(sysUser)

	sql := ""

	if sysUser.Role != "" {
		sql = sql + "role = \"" + sysUser.Role + "\""
		if sysUser.UserName != "" {
			sql = sql + " and user_name LIKE '%" + sysUser.UserName + "%' "
		}
	} else {
		if sysUser.UserName != "" {
			sql = sql + "user_name LIKE '%" + sysUser.UserName + "%' "
		}
	}

	logger.Info(sql)
	err = tx.Model(sysUser).Where(sql).Count(&count).Error
	if err != nil {
		logger.Error(err)
	}
	err = tx.Model(sysUser).Where(sql).Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&sysUsers).Error
	if err != nil {
		logger.Error(err)
		return sysUsers, count, err
	}
	return sysUsers, count, err
}

func DeleteUserById(id int64, tx *gorm.DB) (err error) {

	err = tx.Delete(&SysUser{ID: id}).Error
	if err != nil {
		return err
	}
	return nil
}
