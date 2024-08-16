package global

import (
	"sync"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 全局客户端
var (
	GvaMysqlClient *gorm.DB               //Mysql客户端
	GvaRedisClient redis.UniversalClient //Redis客户端
)

// 健康状态
var (
	Health = true
	HealthLock sync.RWMutex 
)