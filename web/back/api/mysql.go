package api

import (
	"errors"
	"fmt"
	"kehaha-5/k8s_demo/global"
	"kehaha-5/k8s_demo/model"
	"kehaha-5/k8s_demo/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type getParams struct {
	Id int64 `form:"id" binding:"gte=0,required"`
}

type setParams struct {
	UserName  string       `json:"user_name" binding:"required" validate:"required"`
	UserPhone uint         `json:"user_phone" binding:"required" validate:"required,numeric"`
	Age       uint         `json:"age" binding:"required,gte=0,lte=120" validate:"required,gte=0,lte=120"`
	Sex       model.Gender `json:"sex" binding:"required,oneof=1 2" validate:"required,oneof=1 2"` // assuming 0 and 1 represent valid genders
}

func mysql_get(ctx *gin.Context) {
	var getParams getParams
	if err := ctx.ShouldBindQuery(&getParams); err != nil {
		response.ErrorWithMsg(ctx, err.Error())
		return
	}
	var user model.User
	res := global.GvaMysqlClient.First(&user, getParams.Id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			response.ErrorWithMsg(ctx, "用户不存在")
			return
		}
		response.Error(ctx)
		fmt.Printf("error to query:%s\n", res.Error.Error())
		return
	}
	response.OkWithData(ctx, user)
}

func mysql_set(ctx *gin.Context) {
	var setParams setParams
	if err := ctx.ShouldBindJSON(&setParams); err != nil {
		response.ErrorWithMsg(ctx, err.Error())
		return
	}
	var user model.User
	user.Age = setParams.Age
	user.Sex = setParams.Sex
	user.UserPhone = setParams.UserPhone
	user.UserName = &setParams.UserName
	res := global.GvaMysqlClient.Create(&user)
	if res.Error != nil {
		response.Error(ctx)
		fmt.Printf("error to insert:%s\n", res.Error.Error())
		return
	}
	response.OkWithData(ctx, user)
}
