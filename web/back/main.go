package main

import (
	"fmt"
	"kehaha-5/k8s_demo/api"
	"kehaha-5/k8s_demo/config"
	"kehaha-5/k8s_demo/health"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	eg errgroup.Group
)

func main() {
	config.InitConfig()
	r := gin.Default()
	api.InitApiRouter(r)

	httpServer := &http.Server{
		//ip和端口号
		Addr: config.Config.App.Addr,
		//调用的处理器，如为nil会调用http.DefaultServeMux
		Handler: r,
	}

	configWdPath, _ := filepath.Abs(config.Config.App.ConfigFile)
	fmt.Printf("run info: \n app name: %s \n app env: %s\n app config file: %s \n app listen: http://%s \n",
		config.Config.App.Desc,
		config.Config.App.Env,
		configWdPath,
		config.Config.App.Addr,
	)
	eg.Go(func() error {
		return httpServer.ListenAndServe()
	})

	healthApiRouter := gin.New()
	health.InitApiRouter(healthApiRouter)
	healthApiRouter.Use(gin.Recovery())
	healthServer := &http.Server{
		Addr:    ":8080",
		Handler: healthApiRouter,
	}
	eg.Go(func() error {
		return healthServer.ListenAndServe()
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
