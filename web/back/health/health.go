package health

import (
	"kehaha-5/k8s_demo/global"
	"kehaha-5/k8s_demo/response"

	"github.com/gin-gonic/gin"
)

func healthApi(ctx *gin.Context) {
	global.HealthLock.RLock()
	defer global.HealthLock.RUnlock()
	if global.Health {
		response.Ok(ctx)
	} else {
		response.Error(ctx)
	}
}
