package router

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/VihKun/github-api-manager/client"
	"github.com/gin-gonic/gin"
)

const (
	basePath = "api/v1"
)

var (
	// Used for calculating uptime
	start = time.Now()
)

func initRoutes(r *gin.Engine, ghClient *client.GitHubClient) {
	v1 := r.Group(basePath)
	{
		// Health check endpoint
		v1.GET("/health", func(c *gin.Context) {
			status := "OK"
			uptime := time.Since(start).String()

			user, _, err := ghClient.Client.Users.Get(context.Background(), "")

			if err != nil {
				status = "ERROR"
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  status,
					"uptime":  uptime,
					"message": "GitHub API is down",
					"error":   err.Error(),
				})
				os.Exit(1)
			}

			c.JSON(http.StatusOK, gin.H{
				"status":  status,
				"uptime":  uptime,
				"message": "Service Running",
				"user":    user.GetLogin(),
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
