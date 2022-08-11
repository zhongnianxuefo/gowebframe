package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime"
	"os"
	"os/signal"
	"syscall"
)

func InitGin() {
	config := GetConfig()
	db := GetDateBase()
	log := GetLogger()

	port := config.Port

	err := mimeType()
	if err != nil {
		log.Error(err)
	}

	Router := gin.Default()

	Router.POST("/api/login", Login)
	Router.POST("/api/do", DoApi)

	Router.Static("/favicon.ico", "./web/favicon.ico")
	Router.Static("/js", "./web/js")           // web里面的js文件
	Router.Static("/css", "./web/css")         // web里面的css文件
	Router.StaticFile("/", "./web/index.html") // 前端网页入口页面

	addr := fmt.Sprintf(":%d", port)

	go func() {
		err = Router.Run(addr)
		if err != nil {
			log.Error(err)
		}
	}()

	//等待中断信号，以优雅地关闭服务器
	quit := make(chan os.Signal)
	// 可以捕捉除了kill-9的所有中断信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	sqlDB, err := db.DB()
	if err != nil {
		log.Error(err)
	}
	sqlDB.Close()

	log.Info("退出程序")

	return
}

// mime 设置
func mimeType() (err error) {
	err = mime.AddExtensionType(".css", "text/css; charset=utf-8")
	if err != nil {
		return
	}
	return
}
