package api

import (
	"github.com/gin-gonic/gin"
)

// 系统路由
func InitApiRouter(engine *gin.Engine) {
	systemRouter := engine.Group("api")
	{
		systemRouter.GET("mysql_get", mysql_get)
		systemRouter.POST("mysql_set", mysql_set)
		systemRouter.GET("redis_get", redis_get)
		systemRouter.POST("redis_set", redis_set)
		systemRouter.GET("set_un_health", set_un_health)
	}

}
