package main

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	UUID     uuid.UUID // 用户UUID
	UserName string    // 用户登录名
	PassWord string    // 用户登录密码
	Enable   bool      //用户是否有效
}
