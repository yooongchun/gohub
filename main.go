package main

import (
	"flag"
	"fmt"
	"gohub/app/http/middlewares"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/auth"
	"gohub/pkg/config"
	"gohub/pkg/response"

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

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	//init gin instance
	router := gin.New()

	// 初始化 DB
	bootstrap.SetupDB()
	bootstrap.SetupRedis()
	// 初始化路由
	bootstrap.SetupRoute(router)

	router.GET("/test_auth", middlewares.AuthJWT(), func(c *gin.Context) {
		userModel := auth.CurrentUser(c)
		response.Data(c, userModel)
	})

	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
