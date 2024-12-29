package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	basePath = "api/v1"
)

var (
	// Used for calculating uptime
	start = time.Now()
)

func initRoutes(r *gin.Engine) {
	v1 := r.Group(basePath)
	{
		// Health check endpoint
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
				"uptime": time.Since(start).String(),
			})
		})

		// Repository endpoints
		v1.GET("/repos", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "List Repositories",
			})
		})

		v1.POST("/repos", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "Create Repository",
			})
		})

		v1.DELETE("/repos/:name", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "Delete Repository",
			})
		})

		// Pull requests endpoint
		v1.GET("/repos/:name/pulls", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "List Pulls Requests",
			})
		})
	}
}
