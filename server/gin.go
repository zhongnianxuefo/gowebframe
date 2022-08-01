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

type Request struct {
	ApiType     string
	RequestData map[string]interface{}
}

type Response struct {
	Request

	ResponseCode int
	ResponseData map[string]interface{}

	ErrMsg    string
	ErrSource string

	c *gin.Context
}

func NewResponse(c *gin.Context) (r Response, err error) {
	r.c = c
	r.RequestData = make(map[string]interface{})
	r.ResponseData = make(map[string]interface{})
	err = c.ShouldBindJSON(&r.Request)
	return
}

func (r *Response) Ok() {
	r.ResponseCode = 0
	r.c.JSON(http.StatusOK, r)
}

func (r *Response) Error(errMsg, errSource string) {
	r.ResponseCode = 1
	r.ErrMsg = errMsg
	r.ErrSource = errSource
	r.c.JSON(http.StatusOK, r)
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
	r, err := NewResponse(c)
	if err != nil {
		log.Error(err)
	}
	//log.Debug("Login:", r)
	userName, _ := r.RequestData["username"].(string)
	passWord, _ := r.RequestData["password"].(string)
	var users []SysUser
	db.Where(&SysUser{UserName: userName}).Find(&users)
	if len(users) == 0 {
		r.Error("用户名不存在！", "username")
	} else {
		user := users[0]
		pwd := fmt.Sprintf("%x", md5.Sum([]byte(passWord)))
		if pwd == user.PassWord {
			r.Ok()
		} else {
			r.Error("密码错误！", "password")
		}
	}
	return
}
