package router

import (
	"fmt"
	"os"

	"github.com/VihKun/github-api-manager/handler"
	"github.com/gin-gonic/gin"
)

func Init() {
	// Initialize Router
	r := gin.Default()

	// Initialize Client in handler/handler.go
	handler.Init()

	// Initialize Routes
	initRoutes(r)

	// Run API Server
	if err := r.Run(); err != nil {
		fmt.Printf("Error runnig API server: %v", err)
		os.Exit(1)
	}
}
