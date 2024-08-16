package health

import (
	"github.com/gin-gonic/gin"
)

// 系统路由
func InitApiRouter(engine *gin.Engine) {
	systemRouter := engine.Group("api")
	{
		systemRouter.GET("health", healthApi)
	}

}
