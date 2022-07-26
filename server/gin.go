package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mime"
)

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

func ginMain(addr string) {
	err := mimeType()
	if err != nil {
		log.Fatalln(err)
	}

	Router := gin.Default()

	Router.POST("/api", DoApi)
	//Router.GET("/widget/:name", getWidgetJson)
	//Router.GET("/widgetdata/:name", getWidgetData)

	// Router.LoadHTMLGlob("./web/*.html")
	//Router.Static("/favicon.ico", "./web/favicon.ico")
	//Router.Static("/js", "./web/js")           // web里面的js文件
	//Router.Static("/css", "./web/css")         // web里面的css文件
	//Router.StaticFile("/", "./web/index.html") // 前端网页入口页面

	err = Router.Run(addr)
	if err != nil {
		log.Fatalln(err)
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
