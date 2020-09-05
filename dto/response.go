package dto

import "github.com/gin-gonic/gin"

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func CreateResponse(ctx *gin.Context, httpCode int, code int, msg string, data interface{})  {
	ctx.JSON(httpCode, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Success(ctx *gin.Context, data interface{})  {
	CreateResponse(ctx, 200, 200, "请求成功", data)
}

func Failed(ctx *gin.Context, code int, msg string)  {
	CreateResponse(ctx, 200, code, msg, nil)
}
