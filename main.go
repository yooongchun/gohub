package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	//init gin instance
	r := gin.New()

	//register a middware
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		//response  using JSON
		c.JSON(http.StatusOK, gin.H{
			"Hello": "world!",
		})
	})

	// 404 handler
	r.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "Page not found!")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "route not defined!",
			})
		}
	})
	fmt.Println("Hello World! 世界!")
	r.Run(":8080")
}
