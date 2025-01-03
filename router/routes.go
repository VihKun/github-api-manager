package router

import (
	"github.com/VihKun/github-api-manager/handler"
	"github.com/gin-gonic/gin"
)

const (
	basePath = "api/v1"
)

func initRoutes(r *gin.Engine) {
	v1 := r.Group(basePath)
	{
		// Health check endpoint
		v1.GET("/health", handler.HealthCheckHandler)

		// Repository endpoints
		v1.GET("/repos", handler.ListReposHandler)

		v1.POST("/repos", handler.CreateRepoHandler)

		v1.DELETE("/repos/:name", handler.DeleteRepoHandler)

		// Pull requests endpoint
		v1.GET("/repos/:name/pulls", handler.ListPullRequestsHandler)
	}
}
