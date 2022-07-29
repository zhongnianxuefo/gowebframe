package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

//打开数据库
func openDateBase(config SoftwareConfig) (db *gorm.DB, err error) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 5000 * time.Millisecond,
			LogLevel:      logger.Warn,
			Colorful:      true,
		},
	)

	if config.DatabaseType == "postgres" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			config.DatabaseHost, config.DatabaseUser, config.DatabasePassword, config.DatabaseName, config.DatabasePort)

		d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
		if err != nil {
			return nil, err
		}
		db = d
	} else {
		return nil, errors.New("数据库类型错误")
	}

	return
}

func InitDataBase(db *gorm.DB) (err error) {

	err = db.AutoMigrate(&SysUser{})
	if err != nil {
		return
	}

	err = initSysUser(db)
	if err != nil {
		return
	}

	return
}
func initSysUser(db *gorm.DB) (err error) {
	var c int64
	err = db.Model(&SysUser{}).Count(&c).Error
	if err != nil {
		return
	}
	if c == 0 {
		var u SysUser
		u.UUID = uuid.NewV4()
		u.UserName = "admin"
		u.PassWord = fmt.Sprintf("%x", md5.Sum([]byte("admin")))
		u.Enable = true
		db.Save(&u)
	}

	return
}
