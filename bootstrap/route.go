// Package bootstrap: init route

package bootstrap

import (
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// SetupRoute init route
func SetupRoute(router *gin.Engine) {
	// register global middleware
	registerGlobalMiddleWare(router)

	// register API routes
	routes.RegisterAPIRoutes(router)

	// setup 404 router
	setup404Handler(router)

}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery())
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "404: Page not found!")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "Route not defined!",
			})
		}
	})
}
