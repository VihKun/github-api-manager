package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/VihKun/github-api-manager/client"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v50/github"
)

var (
	// Used for the GitHub API client
	ghClient *client.GitHubClient
	// Used for calculating uptime
	start = time.Now()
)

func Init() {
	var err error
	// Initialize Client
	ghClient, err = client.InitClient()
	if err != nil {
		fmt.Printf("Error initializing GitHub client: %v", err)
		os.Exit(1)
	}
}

func HealthCheckHandler(c *gin.Context) {
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
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  status,
		"uptime":  uptime,
		"message": "Service Running",
		"user":    user.GetLogin(),
	})
}

func ListReposHandler(c *gin.Context) {
	// Check if client is initialized
	if ghClient == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "GitHub client is not initialized"})
		return
	}

	// Fetch username
	user, _, err := ghClient.Client.Users.Get(c, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// List the repositories
	repos, _, err := ghClient.Client.Repositories.List(c, *user.Login, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the list of repositories
	c.JSON(http.StatusOK, repos)
}

/*
	{
	  "name": "name_of_repo",
	  "description": "description_of_repo",
	  "private": true/false
	}
*/
func CreateRepoHandler(c *gin.Context) {
	// Parse the request body to get the repo details
	var repo github.Repository
	if err := c.ShouldBindJSON(&repo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Create the repo using the GitHub API client
	createdRepo, _, err := ghClient.Client.Repositories.Create(context.Background(), "", &repo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create repository: %v", err)})
		return
	}

	// Return the created repo
	c.JSON(http.StatusOK, createdRepo)
}

func DeleteRepoHandler(c *gin.Context) {
	// Fetch name of the repository
	repo := c.Param("name")

	// Check if a repo was provided
	if repo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repo field is required"})
		return
	}

	// Fetch username
	user, _, err := ghClient.Client.Users.Get(c, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Try to delete repo
	_, err2 := ghClient.Client.Repositories.Delete(context.Background(), *user.Login, repo)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete repository: %v", err2)})
		return
	}

	// Return Deletion successful
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Repository %s/%s was deleted successfully", *user.Login, repo)})
}

func ListPullRequestsHandler(c *gin.Context) {
	// Fetch name of the repository
	repo := c.Param("name")

	// Check if a repo was provided
	if repo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Owner and repo parameters are required"})
		return
	}

	// Fetch username
	user, _, err := ghClient.Client.Users.Get(c, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Try to get pull requests
	pulls, _, err := ghClient.Client.PullRequests.List(context.Background(), *user.Login, repo, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to list pull requests: %v", err)})
		return
	}

	// Return Pull Requests
	c.JSON(http.StatusOK, pulls)
}
