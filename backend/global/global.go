package global

import (
	"errors"
	"gorm.io/gorm"
)

var (
	CfgFilePath = "./config/config.toml"
	DB          *gorm.DB
)

var (
	ErrCreated  = errors.New("创建失败")
	ErrDeleted  = errors.New("删除失败")
	ErrUpdated  = errors.New("更新失败")
	ErrGet      = errors.New("查询失败")
	ErrRegister = errors.New("注册失败")
	ErrLogin    = errors.New("用户名或密码错误")
	ErrAuth     = errors.New("未登录或非法访问")
)
