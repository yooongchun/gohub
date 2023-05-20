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

	//init gin instance
	router := gin.New()

	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
