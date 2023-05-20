// Register routes
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// register route group
	v1 := r.Group("v1")
	{
		// register root api route
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "world",
			})
		})
	}
}
