package service

import (
	"errors"
	"github.com/nc-77/gtils"
	"gorm.io/gorm"
	"todo/global"
	"todo/model"
)

func Register(user *model.User) (err error) {
	if !errors.Is(global.DB.Where("username = ?", user.Username).First(&model.User{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已被注册")
	}
	user.Password = gtils.MD5V(user.Password)
	if err = global.DB.Create(user).Error; err != nil {
		return global.ErrRegister
	}
	return nil
}

func Login(user *model.User) (err error) {
	user.Password = gtils.MD5V(user.Password)
	if err = global.DB.Where("username = ? and password = ?", user.Username, user.Password).First(user).Error; err != nil {
		return global.ErrLogin
	}
	return nil
}
func FindUserById(id int) (err error, user model.User) {
	if err = global.DB.Where("id = ?", id).Find(&user).Error; err != nil {
		return errors.New("用户不存在"), user
	}
	return
}
