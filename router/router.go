package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Init() {
	// Initialize Router
	r := gin.Default()

	// Initialize Routes
	initRoutes(r)

	// Run API Server
	if err := r.Run(); err != nil {
		fmt.Printf("Error runnig API server: %v", err)
		os.Exit(1)
	}
}
