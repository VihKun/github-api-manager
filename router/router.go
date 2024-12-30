package router

import (
	"fmt"
	"os"

	"github.com/VihKun/github-api-manager/client"
	"github.com/gin-gonic/gin"
)

func Init() {
	// Initialize Router
	r := gin.Default()

	// Initialize Client
	ghClient, err := client.InitClient()
	if err != nil {
		fmt.Printf("Error initializing GitHub client: %v", err)
	}

	// Initialize Routes
	initRoutes(r, ghClient)

	// Run API Server
	if err := r.Run(); err != nil {
		fmt.Printf("Error runnig API server: %v", err)
		os.Exit(1)
	}
}
