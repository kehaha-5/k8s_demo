package config

import (
	"context"
	"flag"
	"fmt"
	"kehaha-5/k8s_demo/global"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var configFile string

/**
*  parseCommand
*  @Desc：解析命令行参数
*  @Author LiuQingHui
*  @Date 2021-11-02 10:31:54
**/
func parseCommand() {
	// 读取配置文件优先级: 命令行 > 默认值
	flag.StringVar(&configFile, "c", "./config.yaml", "配置")
	fmt.Println("configFile:", configFile)
	flag.Parse()
}

// ViperInit 初始化viper配置解析包，函数可接受命令行参数
func InitConfig() {
	parseCommand()
	if len(configFile) == 0 {
		// 读取默认配置文件
		panic("配置文件不存在！")
	}
	// 读取配置文件
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置解析失败:%s", err))
	}
	// 动态监测配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生改变")
		if err := v.Unmarshal(&Config); err != nil {
			panic(fmt.Errorf("配置重载失败:%s", err))
		}
	})
	if err := v.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("配置重载失败:%s", err))
	}
	// 设置配置文件
	Config.App.ConfigFile = configFile

	initMysql()
	initRedis()
}

func initMysql() {
	mysqlConfig := Config.Mysql
	if !mysqlConfig.Enable {
		return
	}
	masterDns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.MasterHost, mysqlConfig.Port, mysqlConfig.Database, mysqlConfig.Charset,
		mysqlConfig.ParseTime, mysqlConfig.TimeZone)

	client, err := gorm.Open(mysql.Open(masterDns), &gorm.Config{
		SkipDefaultTransaction: mysqlConfig.Gorm.SkipDefaultTx, //是否跳过默认事务
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   mysqlConfig.Gorm.TablePrefix,
			SingularTable: mysqlConfig.Gorm.SingularTable,
		},
		// 执行任何SQL时都会创建一个prepared statement并将其缓存，以提高后续的效率
		PrepareStmt: mysqlConfig.Gorm.PreparedStmt,
		//在AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true
		DisableForeignKeyConstraintWhenMigrating: mysqlConfig.Gorm.CloseForeignKey,
	})

	if err != nil {
		panic(fmt.Sprintf("创建mysql客户端失败: %s, %v", err, mysqlConfig))
	}

	if len(mysqlConfig.SlaverHost) != 0 {
		replacerDns := []gorm.Dialector{}
		for _, it := range mysqlConfig.SlaverHost {
			slaver := mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
				mysqlConfig.User, mysqlConfig.Password, it, mysqlConfig.Port, mysqlConfig.Database, mysqlConfig.Charset,
				mysqlConfig.ParseTime, mysqlConfig.TimeZone))
			replacerDns = append(replacerDns, slaver)
		}
		err := client.Use(dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(masterDns)},
			Replicas: replacerDns,
			// sources/replicas 负载均衡策略
			Policy: dbresolver.RandomPolicy{},
		}))

		if err != nil {
			fmt.Printf("开启从库配置成功")
		} else {
			panic(fmt.Sprintf("开启从库配置失败: %s, %v", err, replacerDns))
		}
	}

	// 赋值给全局变量
	global.GvaMysqlClient = client

}

func initRedis() {
	if !Config.Redis.Enable {
		return
	}
	// 创建
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:       Config.Redis.Addrs,
		Password:    Config.Redis.Password,
		DB:          Config.Redis.DefaultDB,
		DialTimeout: Config.Redis.DialTimeout,
	})

	// 使用超时上下文，验证redis
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), Config.Redis.DialTimeout)
	defer cancelFunc()
	_, err := redisClient.Ping(timeoutCtx).Result()
	if err != nil {
		panic("redis初始化失败! " + err.Error())
	}
	global.GvaRedisClient = redisClient
}
