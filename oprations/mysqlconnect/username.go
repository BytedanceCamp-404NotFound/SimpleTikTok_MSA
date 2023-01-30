package mysqlconnect

import (
	"fmt"
	"time"

	"SimpleTikTok/tools/encryption"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// 向user_login表中插入新值，同时更新user表
// 返回id不为-1表示注册成功
// 返回id为-1表示失败
func CreateUser(db *gorm.DB, UserName string, password string) int {
	id := -1
	// db, _ := SqlConnect()
	user := User_login{UserName: UserName, UserPwd: encryption.HashEncode(password), RegisterDate: time.Now()}
	db.Table("user_login").Create(&user)
	db.Table("user_login").Select("user_id").Where("user_name = ? and user_pwd = ?", UserName, encryption.HashEncode(password)).Find(&id)

	CreateInfo(UserName, int64(id))
	return id
}

// 函数功能 校验user_name表中的账户密码是否一致
// 返回id不为-1表示一致
// 返回id为-1表示不一致
func CheckUser(UserName string, password string) (int, error) {
	id := -1
	db := GormDB
	err := db.Table("user_login").Select("user_id").Where("user_name = ? and user_pwd = ?", UserName, encryption.HashEncode(password)).Find(&id).Error
	if err != nil {
		logx.Errorf("Check user fail, error:%v", err.Error())
		return -1, err
	}
	fmt.Println(id)

	return id, err
}

// 检查注册用户是否存在,存在则修改密码,不存在才创建用户
// -1代表发生错误
// 0代表用户不存在
// 否则代表用户存在
func FindUserIsExist(db *gorm.DB, userName string, password string) (int, error) {
	var count int64
	var uid int
	// db, _ := SqlConnect()
	err := db.Table("user_login").Select("user_id").Where("user_name = ?", userName).Find(&uid).Count(&count).Error
	if err != nil {
		return -1, err
	}
	if count == 0 {
		return 0, nil
	}
	return uid, nil
}
