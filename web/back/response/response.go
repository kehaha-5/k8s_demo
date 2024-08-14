package response

import (
	"kehaha-5/k8s_demo/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 200
	ERROR   = 500
)

// 定义统一返回接口格式
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time string      `json:"-"`
}

// 请求响应
func ResultJson(ctx *gin.Context, code int, msg string, data interface{}) {
	// 格式化时间
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
		Time: time.Now().Format(config.Config.App.TimeFormat),
	})
}

// 返回固定成功信息
func Ok(ctx *gin.Context) {
	ResultJson(ctx, SUCCESS, "请求成功", map[string]interface{}{})
}

// 只返回成功消息
func OkWithMsg(ctx *gin.Context, msg string) {
	ResultJson(ctx, SUCCESS, msg, map[string]interface{}{})
}

// 返回固定消息和数据
func OkWithData(ctx *gin.Context, data interface{}) {
	ResultJson(ctx, SUCCESS, "请求成功", data)
}

// 返回指定消息和数据
func OkWithDetail(ctx *gin.Context, msg string, data interface{}) {
	ResultJson(ctx, SUCCESS, msg, data)
}

// 错误信息
func ErrorWithMsg(ctx *gin.Context, msg string) {
	ResultJson(ctx, ERROR, msg, map[string]interface{}{})
}

func Error(ctx *gin.Context) {
	ResultJson(ctx, ERROR, "服务器错误", map[string]interface{}{})
}
