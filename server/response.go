package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Request struct {
	Token       string
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

const (
	ResponseCodeOk = iota
	ResponseCodeErr
	ResponseCodeTokenErr
)

func NewResponse(c *gin.Context) (r Response, err error) {
	r.c = c

	r.RequestData = make(map[string]interface{})
	r.ResponseData = make(map[string]interface{})
	err = c.ShouldBindJSON(&r.Request)
	return
}

func (r *Response) getUserName() (userName string, err error) {
	if r.Token == "" {
		r.TokenError("token不能为空")
		r.c.Abort()
		return
	}
	mc, err := ParseToken(r.Token)
	if err != nil {
		r.TokenError(err.Error())
		r.c.Abort()
		return
	}
	userName = mc.UserName
	return
}

func (r *Response) Ok() {
	r.ResponseCode = ResponseCodeOk
	r.c.JSON(http.StatusOK, r)
}
func (r *Response) OkWithData(data map[string]interface{}) {
	r.ResponseCode = ResponseCodeOk
	r.ResponseData = data
	r.c.JSON(http.StatusOK, r)
}
func (r *Response) Error(errMsg, errSource string) {
	r.ResponseCode = ResponseCodeErr
	r.ErrMsg = errMsg
	r.ErrSource = errSource
	r.c.JSON(http.StatusOK, r)
}

func (r *Response) TokenError(errMsg string) {
	r.ResponseCode = ResponseCodeTokenErr
	r.ErrMsg = errMsg
	r.c.JSON(http.StatusOK, r)
}
