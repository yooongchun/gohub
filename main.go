package main

import (
	"flag"
	"fmt"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}
func main() {
	// load env
	var env string
	flag.StringVar(&env, "env", "", "load .env file, using .env.testing when \"--env=testing\" used")
	flag.Parse()
	config.InitConfig(env)

	// 初始化 Logger
	bootstrap.SetupLogger()

	//init gin instance
	router := gin.New()

	// 初始化 DB
	bootstrap.SetupDB()
	// 初始化路由
	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
