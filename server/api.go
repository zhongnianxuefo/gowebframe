package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

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
			r.Token, err = GetToken(user.UserName)
			if err != nil {
				log.Error(err)
			}
			r.Ok()
		} else {
			r.Error("密码错误！", "password")
		}
	}
	return
}

func DoApi(c *gin.Context) {
	log := GetLogger()
	r, err := NewResponse(c)
	if err != nil {
		log.Error(err)
	}
	userName, err := r.getUserName()
	if err != nil {
		log.Error(err)
	}
	if userName == "" {
		return
	}
	if r.ApiType == "test" {
		r.Ok()
	} else if r.ApiType == "getAllMenus" {
		getAllMenus(r)
	}

}

func getAllMenus(r Response) {
	log := GetLogger()

	jsonData := `
	{
		"AllMenus":[
		  {"title":"首页", "icon":"home-outlined", "page":"home"},
		  {"title":"基础资料", "icon":"user-outlined", "subMenus":[
			  {"title":"客户信息", "icon":"team-outlined", "page":"list","module":"客户信息"},
			  {"title":"工厂信息",  "page":"list","module":"工厂信息"},
			  {"title":"产品信息", "icon":"database-outlined", "page":"list","module":"产品信息"}
			]},
		  {"title":"业务相关",  "subMenus":[
			  {"title":"外销合同",  "page":"list","module":"外销合同"},
			  {"title":"采购合同",  "page":"list","module":"采购合同"}
			]}
		]
	}
	`
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Error(err)
	}

	r.OkWithData(data)
}
