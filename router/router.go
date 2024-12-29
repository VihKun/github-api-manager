package router

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	// Used for calculating uptime
	start = time.Now()
)

func Init() {
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
			"uptime": time.Since(start).String(),
		})
	})

	// Run API Server
	if err := r.Run(); err != nil {
		fmt.Printf("Error runnig API server: %v", err)
		os.Exit(1)
	}
}
