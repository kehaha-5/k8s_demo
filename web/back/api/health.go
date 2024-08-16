package api

import (
	"kehaha-5/k8s_demo/global"
	"kehaha-5/k8s_demo/response"

	"github.com/gin-gonic/gin"
)

func set_un_health(ctx *gin.Context) {
	global.HealthLock.Lock()
	defer global.HealthLock.Unlock()
	global.Health = false
	response.Ok(ctx)
}
