package api

import (
	"fmt"
	"kehaha-5/k8s_demo/global"
	"kehaha-5/k8s_demo/response"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type getRedisParam struct {
	Key string `form:"key" binding:"required"`
}

type setRedisParam struct {
	Key   string      `form:"key" binding:"required"`
	Value interface{} `form:"value" binding:"required"`
}

func redis_get(ctx *gin.Context) {
	var getParams getRedisParam
	if err := ctx.ShouldBindQuery(&getParams); err != nil {
		response.ErrorWithMsg(ctx, err.Error())
		return
	}
	res := global.GvaRedisClient.Get(ctx, getParams.Key)
	switch {
	case res.Err() == redis.Nil:
		response.ErrorWithMsg(ctx, "值不存在")
	case res.Err() != nil:
		response.Error(ctx)
		fmt.Printf("error to query:%s\n", res.Err().Error())
	default:
		response.OkWithData(ctx, res.Val())
	}
}

func redis_set(ctx *gin.Context) {
	var setParams setRedisParam
	if err := ctx.ShouldBindJSON(&setParams); err != nil {
		response.ErrorWithMsg(ctx, err.Error())
		return
	}
	res := global.GvaRedisClient.Set(ctx, setParams.Key, setParams.Value, 0)
	if res.Err() != nil {
		response.Error(ctx)
		fmt.Printf("error to set:%s\n", res.Err().Error())
		return
	}
	data := make(map[string]interface{})
	data["key"] = setParams.Key
	data["value"] = res.Val()
	response.OkWithData(ctx, data)
}
