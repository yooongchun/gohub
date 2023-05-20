package main

import (
	"fmt"
	"gohub/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	//init gin instance
	router := gin.New()

	bootstrap.SetupRoute(router)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}
