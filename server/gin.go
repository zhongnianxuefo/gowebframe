package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Response struct {
	Code int
	Data interface{}
	Msg  string
}

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

	Router.POST("/api/do", DoApi)
	Router.POST("/api/login", Login)

	//Router.GET("/widget/:name", getWidgetJson)
	//Router.GET("/widgetdata/:name", getWidgetData)

	// Router.LoadHTMLGlob("./web/*.html")
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

//显示页面错误信息
func showHtmlError(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"success": false,
		"error":   err.Error(),
	})
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

func DoApi(c *gin.Context) {

	doInfo := make(map[string]interface{})
	_ = c.ShouldBindJSON(&doInfo)

	s, _ := json.Marshal(doInfo)
	fmt.Println("doInfo :")
	fmt.Println(string(s))
}

func Login(c *gin.Context) {
	db := GetDateBase()
	log := GetLogger()

	doInfo := make(map[string]interface{})
	log.Debug("Login:", doInfo)

	err := c.ShouldBindJSON(&doInfo)
	if err != nil {
		log.Error(err)
	}

	userName := doInfo["username"].(string)
	passWord := doInfo["password"].(string)

	var u SysUser
	var count int64
	db.Find(&SysUser{UserName: userName}).Count(&count)
	if count == 0 {
		responseErr(c, doInfo, "用户名不存在！")
	} else {
		db.Find(&SysUser{UserName: userName}).First(&u)
		pwd := fmt.Sprintf("%x", md5.Sum([]byte(passWord)))
		if pwd == u.PassWord {
			responseOk(c, doInfo)
		} else {
			responseErr(c, doInfo, "密码错误！")
		}
	}

	//s, _ := json.Marshal(doInfo)
	//fmt.Println("Login :")
	//fmt.Println(string(s))

	//Global.Logger.Debug("Login :", doInfo)
	//responseOk(c, doInfo)
	return

}
func responseOk(c *gin.Context, data map[string]interface{}) {
	data["返回代码"] = 0
	c.JSON(http.StatusOK, data)

}

func responseErr(c *gin.Context, data map[string]interface{}, err string) {
	data["返回代码"] = 1
	data["返回错误信息"] = err
	c.JSON(http.StatusOK, data)

}
